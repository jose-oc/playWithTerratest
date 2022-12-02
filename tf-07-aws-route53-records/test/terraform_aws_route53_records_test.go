package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/random"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// Zone name defined in your AWS account
const zoneName = "myzonename.io"
const randomNameKey = "randomName"

func TestTerraformAwsRoute53Records(t *testing.T) {
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
	randomName := fmt.Sprintf("datacenter_%s", strings.ToLower(random.UniqueId()))
	test_structure.SaveString(t, workingDir, randomNameKey, randomName)

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: workingDir,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"create":       true,
			"private_zone": false,
			"zone_name":    zoneName,

			"records": []map[string]interface{}{
				{
					"name":            "esxi01." + randomName,
					"type":            "A",
					"ttl":             3600,
					"records":         []string{"70.187.70.91"},
					"allow_overwrite": true,
				},
				{
					"name":            "esxi02." + randomName,
					"type":            "A",
					"ttl":             3600,
					"records":         []string{"70.187.70.92"},
					"allow_overwrite": true,
				},
				{
					"name":            "esxi03." + randomName,
					"type":            "A",
					"ttl":             3600,
					"records":         []string{"70.187.70.93"},
					"allow_overwrite": true,
				},
			},
		},
	})

	// Save the Terraform Options struct, instance name, and instance text so future test stages can use it
	test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}

func validateInstanceRunningWebServer(t *testing.T, workingDir string) {
	subdomainName := test_structure.LoadString(t, workingDir, randomNameKey)
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	// Run `terraform output` to get the IP of the instance
	route53RecordFQDN := terraform.OutputMap(t, terraformOptions, "route53_record_fqdn")
	assert.Len(t, route53RecordFQDN, 3)
	expectedRoute53RecordFQDN := map[string]string{
		"esxi01." + subdomainName + " A": "esxi01." + subdomainName + "." + zoneName,
		"esxi02." + subdomainName + " A": "esxi02." + subdomainName + "." + zoneName,
		"esxi03." + subdomainName + " A": "esxi03." + subdomainName + "." + zoneName,
	}
	assert.EqualValues(t, expectedRoute53RecordFQDN, route53RecordFQDN)
}

// Undeploy the terraform-packer-example using Terraform
func undeployUsingTerraform(t *testing.T, workingDir string) {
	// Load the Terraform Options saved by the earlier deploy_terraform stage
	terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)

	terraform.Destroy(t, terraformOptions)
}
