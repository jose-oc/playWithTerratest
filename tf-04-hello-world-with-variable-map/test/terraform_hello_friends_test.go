package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloRockBand(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "..",
		Vars: map[string]interface{}{
			"your_name": "Mr. Robot",
			"rock_band": map[string]string{"drums": "Ana", "bass": "Pepa", "guitar": "Alicia"},
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	countRockBand := terraform.Output(t, terraformOptions, "count_rock_band")
	assert.Equal(t, "Hello Mr. Robot, there are 3 members in your rock band.", countRockBand)
	// Notice the type of output, in this case we get a list. Using `terraform.Output` we get a string not a list
	rockBand := terraform.OutputMap(t, terraformOptions, "rock_band")
	assert.Len(t, rockBand, 3)
}
