<br>
<div align="center">

  <!-- PROJECT LOGO -->

  <a href="https://registry.terraform.io/">
    <img alt="Terraform" src="https://github.com/Nerdware-LLC/.github/blob/main/org_assets/terraform_banner.png" width="600px" />
  </a>

  <!-- PROJECT NAME/HEADER -->

  <h1>Terraform Module Template Repo</h1>

  <!-- PROJECT TAGLINE -->

**🚀 An Awesome Template to Jumpstart Terraform Modules 🚀**

  <!-- PROJECT SHIELDS -->

[![pre-commit][pre-commit-shield]](https://github.com/pre-commit/pre-commit)
[![semantic-release][semantic-shield]](https://github.com/semantic-release/semantic-release)
[![license][license-shield]](/LICENSE)

</div>

---

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
<!-- prettier-ignore-start -->

---

## ⚙️ Module Usage

### Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | 1.2.2 |

### Providers

No providers.

### Modules

No modules.

### Resources

No resources.

### Inputs

No inputs.

### Outputs

No outputs.

---

<!-- prettier-ignore-end -->

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## 🧪 Testing

This module's test suite is powered by [**Terratest**][terratest-url] and can be configured via the file [**tests/test_params.yaml**][test-params-file], which is used to define which assertions the test should make, and what values to use in those assertions.

Input variables to be used for testing purposes can be specified in the file [**tests/test.tfvars**](/tests/test.tfvars). _Note that by default, this file will not be ignored by git, so sensitive values should not be included._ To override this behavior, remove the entry "!tests/test.tfvars" from the [.gitignore](/.gitignore).

#### **Currently Supported Test Params:**

- `expect_outputs` Map the names of expected outputs to arbitrary values of any type. The test will compare these expected outputs against the module's actual outputs; the test will fail if an expected name/key is not present, or if the expected value does not match the actual output value.
- `verbose_logging` Increase test logging verbosity; useful for debugging purposes, or when running tests locally. May not be suitable for some CI environments.

> **_Coming Soon:_** Future releases will expand the supported test params to provide additional functionality. Ideas currently in the pipeline are listed below - PRs are welcome if you've something to add.
>
> - `expect_http_response` Send configurable HTTP requests to created infrastructure components in order for assertions to be made using the resultant responses.
> - `expect_error` Make configurable error assertions.

#### **Test Process Outline:**

1. The [**test_params file**][test-params-file] is read into memory. If invalid test params were provided, the test will fail. If the test params are valid, they're printed to stdout.
2. Your module files are passed into [**Terratest**][terratest-url], which is used to run `terraform init`, `terraform apply`, and `terraform output` to retrieve module outputs.
3. The provided test params are used to determine which assertions to test. For example, providing **expect_outputs** will cause the actual output keys and values to be compared against the expected outputs.
4. Finally, once all assertions have been tested, [**Terratest**][terratest-url] is used to run `terraform destroy` to clean up all test resources - regardless of whether the tests succeed or fail.

<br>
<details>
  <summary><b>Example Test Failure Output</b></summary><br>

If a test fails, the output will look something like the example below. Note: if verbose_logging is not enabled, the first four lines would be omitted.

```shell
TestTerraformModule 2022-06-14T12:35:17-04:00 module_unit_test.go:98:

        [TEST] EXPECTED OUTPUT: foo_string_key = foo_string_value        ACTUAL: foo_string_value_NOPE

module_unit_test.go:91:
                Error Trace:    module_unit_test.go:91
                Error:          Not equal:
                                expected: "foo_string_value"
                                actual  : "foo_string_value_NOPE"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -foo_string_value
                                +foo_string_value_NOPE
                Test:           TestTerraformModule
```

</details>

## 📝 License

Apache 2 Licensed. See [LICENSE](/LICENSE) for more information.

<div align="center" style="margin-top:35px;">

## 💬 Template Author Contact

Trevor Anderson - [@TeeRevTweets](https://twitter.com/teerevtweets) - [T.AndersonProperty@gmail.com](mailto:T.AndersonProperty@gmail.com)

  <a href="https://www.youtube.com/channel/UCguSCK_j1obMVXvv-DUS3ng">
    <img src="https://github.com/trevor-anderson/trevor-anderson/blob/main/assets/YouTube_icon_circle.svg" height="40" />
  </a>
  &nbsp;
  <a href="https://www.linkedin.com/in/trevor-anderson-3a3b0392/">
    <img src="https://github.com/trevor-anderson/trevor-anderson/blob/main/assets/LinkedIn_icon_circle.svg" height="40" />
  </a>
  &nbsp;
  <a href="https://twitter.com/TeeRevTweets">
    <img src="https://github.com/trevor-anderson/trevor-anderson/blob/main/assets/Twitter_icon_circle.svg" height="40" />
  </a>
  &nbsp;
  <a href="mailto:T.AndersonProperty@gmail.com">
    <img src="https://github.com/trevor-anderson/trevor-anderson/blob/main/assets/email_icon_circle.svg" height="40" />
  </a>
  <br><br>

  <a href="https://daremightythings.co/">
    <strong><i>Dare Mighty Things.</i></strong>
  </a>

</div>

<!-- LINKS -->

[pre-commit-shield]: https://img.shields.io/badge/pre--commit-33A532.svg?logo=pre-commit&logoColor=F8B424&labelColor=gray
[semantic-shield]: https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-E10079.svg
[semantic-gh-action-url]: https://github.com/cycjimmy/semantic-release-action
[license-shield]: https://img.shields.io/badge/license-Apache_2.0-000080.svg?labelColor=gray
[terratest-url]: https://terratest.gruntwork.io/docs/
[test-params-file]: /tests/test_params.yaml
