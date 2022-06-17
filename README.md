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

<!-- TODO remove the below section upon completion of post-init setup -->

### **_Post-Initialization Setup_**

1. Review the base template files to ensure your project's needs are met. Some helpful tips:

   - If you'd like to change the [.gitignore](/.gitignore), GitHub has some [awesome templates here](https://github.com/github/gitignore).
   - You can also query the [gitignore.io API](https://docs.gitignore.io/install/command-line) to find a list of recommended gitignore entries to suit virtually any type of project.

     ```bash
     # Obtain a list of available project-type options from the gitignore.io API.
     curl -sL https://www.toptal.com/developers/gitignore/api/list | sed 's/,/\n/g' > ./gitignore_io_api_options

     # Review the resultant list in "./gitignore_io_api_options" to find options that fit your project.
     # You can query gitignore entries for one or more options by separating them with commas.
     # For example, if your project will contain both Terraform and Terragrunt files:
     curl -sL https://www.toptal.com/developers/gitignore/api/terraform,terragrunt >> .gitignore
     ```

2. Set up [**pre-commit**](https://pre-commit.com/#install):

   1. Ensure it's [installed](https://pre-commit.com/#install) locally or in an executable image.
   2. Some awesome pre-commit hooks have already been added to the [**pre-commit config file**](/.pre-commit-config.yaml) for Terraform projects. If you're looking for more hooks to add, the pre-commit project provides a complete list of [supported hooks here](https://pre-commit.com/hooks.html). Some popular hook sources:
      - ["Out-of-the-Box" pre-commit Hooks](https://github.com/pre-commit/pre-commit-hooks)
      - [pre-commit Hooks from gruntwork.io](https://github.com/gruntwork-io/pre-commit)
      - [Some Terraform-specific pre-commit Hooks](https://github.com/antonbabenko/pre-commit-terraform)
   3. Run `pre-commit install` to ensure your git hooks are present.

3. Enable the [**Semantic-Release GitHub Action**][semantic-gh-action-url]:

   1. [Create a GitHub Personal Access Token][gh-pat-docs-url]. When creating the token, the minimum required scopes are:
      - `repo` for a private repository
      - `public_repo` for a public repository
   2. Add a [GitHub Secret][gh-action-docs-url] to your repo named "SEMANTIC_RELEASE_TOKEN" with the value set to the new PAT you created in the previous step.
   3. Once the secret has been added to your repo, you can delete the "check-required-secret" job in the ["Release" GitHub Action workflow](/.github/workflows/release.yaml) (it was included so you can push initialization commits without triggering a bunch of failed GH Action runs). Note that the "Release" workflow will not run unless the ["Test" workflow](/.github/workflows/test.yaml) completes successfully.
   4. (Optional) You can have GH Issues auto-assigned on release failures by adding **assignees** to the "@semantic-release/github" plugin in [.releaserc.yaml](/.releaserc.yaml).

4. Enable [**Terratest**][terratest-url] testing suite:

   1. From the [tests dir](/tests/), run `go mod init "<MODULE_NAME>"` to initialize a go module, where `<MODULE_NAME>` is the name of your module, typically in the format `github.com/<YOUR_USERNAME>/<YOUR_REPO_NAME>`.
      > Example: `go mod init github.com/trevor-anderson/template-terraform-module/tests`.
   2. Run `go mod tidy` to ensure all required package dependencies are installed.
   3. The existing test can be configured using the file [tests/test_params.yaml][test-params-file]. See [**Testing**](#testing) for more info.
      > Your tests will be run by the ["Test" GitHub Action workflow](/.github/workflows/test.yaml), but you can also run them manually from the [tests dir](/tests/) using `go test -v -timeout 30m` (the flags are optional).
   4. In order for Terratest to run within a CI environment, you will need to provision access to your cloud resources in the ["Test" workflow](/.github/workflows/test.yaml). This is best accomplished by [adding GitHub as an OpenID Connect Identity Provider][github-oidc-info-url] within your cloud platform account. If you're operating within a multi-account organization, use a "Sandbox" account.
   5. Once GitHub has been added as an OIDC Identity Provider, add the relevant action for your cloud provider which initiates the OIDC auth flow to authenticate the GitHub OIDC Provider and permits access to the desired cloud resources:

      - AWS: [aws-actions/configure-aws-credentials][github-oidc-aws]
      - Azure: [azure/login][github-oidc-azure]
      - GCP: [google-github-actions/auth][github-oidc-gcp]
      - HashiCorp Vault: [hashicorp/vault-action][github-oidc-vault]
      - Others: [custom action using JWTs][github-oidc-others-custom-jwt]

5. Profit 💰💰💰🥳🎉 _(Also, don't forget to remove this section from the README)_ <!-- https://knowyourmeme.com/memes/profit -->

> Need help? 🤔 Check out my new [**YouTube Channel**](https://www.youtube.com/channel/UCguSCK_j1obMVXvv-DUS3ng) with helpful guides covering how to **_code your cloud_** with Terraform, Terragrunt, Packer, Golang, CI/CD tools, and more.

<!-- Don't remove the pre-commit TF-Docs hook comments, they're used to auto-gen your module's documentation. -->

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
[gh-action-docs-url]: https://docs.github.com/en/actions/security-guides/encrypted-secrets
[gh-pat-docs-url]: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
[terratest-url]: https://terratest.gruntwork.io/docs/
[test-params-file]: /tests/test_params.yaml
[github-oidc-info-url]: https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect
[github-oidc-aws]: https://github.com/marketplace/actions/configure-aws-credentials-action-for-github-actions
[github-oidc-azure]: https://github.com/Azure/login
[github-oidc-gcp]: https://github.com/google-github-actions/auth
[github-oidc-vault]: https://github.com/hashicorp/vault-action
[github-oidc-others-custom-jwt]: https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-cloud-providers
