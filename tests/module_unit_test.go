package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// Structure of the test_params file
type TestParams struct {
	ExpectedOutputs map[string]interface{} `yaml:"expect_outputs,omitempty"`
	VerboseLogging  bool                   `yaml:"verbose_logging,omitempty"`
}

// Test-related file paths (used to obtain absolute paths)
const TEST_PARAMS_FILE = "./test_params.yaml"
const TEST_VARS_FILE = "./test.tfvars"

func TestTerraformModule(t *testing.T) {
	t.Parallel()

	// Get absolute path to test_params file
	testParamsPath, _ := filepath.Abs(TEST_PARAMS_FILE)

	// Get expected outputs from test_params.yaml
	yamlFile, readFileErr := ioutil.ReadFile(testParamsPath)

	// Check for ReadFile err
	if readFileErr != nil {
		logErrAndExit(t, "ReadFile", readFileErr)
	}

	// Initialize test params
	testParams := TestParams{}

	// Unmarshal yaml to retrieve user's test params
	yamlErr := yaml.Unmarshal(yamlFile, &testParams)

	// Check for yaml err
	if yamlErr != nil {
		logErrAndExit(t, "yaml", yamlErr)
	}

	// If verbose logging is enabled, print unmarshaled data to user
	if testParams.VerboseLogging == true {

		// Since it's a struct, we'll need reflection to loop keys/values
		v := reflect.ValueOf(testParams)
		typeOfS := v.Type()

		// Multiline string prefix
		testParamsLogStr := "\n\n\t[TEST] PARAMS: \n"

		// Loop num fields, concat formatted key+value str
		for i := 0; i < v.NumField(); i++ {
			testParamsLogStr += fmt.Sprintf("\t\t- %15s: %v \n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}

		// Print log str
		logger.Log(t, testParamsLogStr)
	}

	// Get absolute path to test.tfvars file
	testVarsPath, _ := filepath.Abs(TEST_VARS_FILE)

	// TF config
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Path to the TF code to be tested
		TerraformDir: "../",
		VarFiles:     []string{testVarsPath},
		// If you need to test Terragrunt configs, use "TerraformBinary": "/path/to/terragrunt"
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs via `terraform output` to compare against expected values.
	actualOutputs := terraform.OutputAll(t, terraformOptions)

	// Loop over ExpectedOutputs, test if actual output equals expected output
	for outputKey, expectedOutputValue := range testParams.ExpectedOutputs {

		actualOutputValue := actualOutputs[outputKey]

		// If verbose logging is enabled, print values before assertion
		if testParams.VerboseLogging {
			logger.Log(t, fmt.Sprintf(
				"\n\n\t[TEST] EXPECTED OUTPUT: %s = %v \t ACTUAL: %v\n",
				outputKey,
				expectedOutputValue,
				actualOutputValue))
		}

		// Verify we're getting back the outputs we expect
		assert.Equal(t, expectedOutputValue, actualOutputValue)
	}
}

func logErrAndExit(t *testing.T, errLabel string, errPayload any) {
	logger.Log(t, fmt.Sprintf("\n\n\t%s ERROR: %v \n\t(EXIT 1) \n", errLabel, errPayload))
	os.Exit(1)
}

/* NOTE re: Logging in CI Environments

Many CI systems will kill your tests if they don't see any log output for a certain
period of time (e.g., 10 minutes in CircleCI). If you use Go's t.Log and t.Logf for
logging in your tests, you'll find that these functions buffer all log output until
the very end of the test (see https://github.com/golang/go/issues/24929 for more info).
If you have a long-running test, this might mean you get no log output for more than
10 minutes, and the CI system will shut down your tests. Moreover, if your test has a
bug that causes it to hang, you won't see any log output at all to help you debug it.

Therefore, we recommend instead using Terratest's logger.Log and logger.Logf functions,
which log to stdout immediately:

	func TestFoo(t *testing.T) {
		logger.Log(t, "This will show up in stdout immediately")
	}

https://terratest.gruntwork.io/docs/testing-best-practices/timeouts-and-logging/
*/
