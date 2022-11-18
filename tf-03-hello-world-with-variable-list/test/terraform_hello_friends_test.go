package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloFriends(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "..",
		Vars: map[string]interface{}{
			"your_name": "Mr. Robot",
			"friends":   []string{"Lis", "Pete", "Cid"},
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	countFriends := terraform.Output(t, terraformOptions, "count_friends")
	assert.Equal(t, "Hello Mr. Robot, you have 3 friends.", countFriends)
	// Notice the type of output, in this case we get a list. Using `terraform.Output` we get a string not a list
	greetingFriends := terraform.OutputList(t, terraformOptions, "greeting_friends")
	assert.Len(t, greetingFriends, 3)
	assert.ElementsMatch(t, greetingFriends, []string{"Hello, Pete!", "Hello, Cid!", "Hello, Lis!"})
}
