package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

const randomNameKey = "randomName"

func TestTerraformAwsStaticWeb(t *testing.T) {
	t.Parallel()

	// The folder where we have our Terraform code
	workingDir := ".."

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer test_structure.RunTestStage(t, "cleanup_terraform", func() {
		undeployUsingTerraform(t, workingDir)
	})

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	// Deploy the web app using Terraform
	test_structure.RunTestStage(t, "deploy_terraform", func() {
		deployUsingTerraform(t, workingDir)
	})

	// Validate that the web app deployed and is responding to HTTP requests
	test_structure.RunTestStage(t, "validate", func() {
		validateInstanceRunningWebServer(t, workingDir)
	})
}

// Deploy the terraform-packer-example using Terraform
func deployUsingTerraform(t *testing.T, workingDir string) {
	randomName := fmt.Sprintf("test.%s", strings.ToLower(random.UniqueId()))
	test_structure.SaveString(t, workingDir, randomNameKey, randomName)

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	awsRegion := aws.GetRandomStableRegion(t, nil, []string{"eu-west-1"})
	t.Logf("We are going to test in region %v", awsRegion)

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: workingDir,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"aws_account_id":     "012345678900",
			"aws_region":         awsRegion,
			"zone_name":          "my-testing-zone.com",
			"private_zone":       false,
			"record_prefix_name": randomName,
		},
	})

	// Save the Terraform Options struct, instance name, and instance text so future test stages can use it
	test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}

func validateInstanceRunningWebServer(t *testing.T, workingDir string) {
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	// Run `terraform output` to get the IP of the instance
	publicIp := terraform.Output(t, terraformOptions, "public_ip")
	privateDns := terraform.Output(t, terraformOptions, "private_dns")
	webserverDnsUrl := terraform.Output(t, terraformOptions, "webserver_dns_url")

	// Make an HTTP request to the instance and make sure we get back a 200 OK with the body "Hello, World!"
	urlIp := fmt.Sprintf("http://%s:80", publicIp)
	urlDns := fmt.Sprintf("http://%s:80", webserverDnsUrl)
	body := fmt.Sprintf("Hello World from %s", privateDns)
	http_helper.HttpGetWithRetry(t, urlIp, nil, 200, body, 30, 5*time.Second)
	http_helper.HttpGetWithRetry(t, urlDns, nil, 200, body, 30, 5*time.Second)
}

// Undeploy the terraform-packer-example using Terraform
func undeployUsingTerraform(t *testing.T, workingDir string) {
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	terraform.Destroy(t, terraformOptions)
}
