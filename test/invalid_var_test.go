package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// simple 'guard-rail' tests
func TestInvalidValues(t *testing.T) {
	inputs := []struct {
		Name        string
		Environment string
		CpuCores    int
		Ram         string
	}{
		{
			Name:        "T0",
			Environment: "sit", // fail
			CpuCores:    4,
			Ram:         "256",
		},
		{
			Name:        "T1",
			Environment: "dev",
			CpuCores:    0, // fail
			Ram:         "256",
		},
		{
			Name:        "T2",
			Environment: "uat",
			CpuCores:    2,
			Ram:         "255", // fail
		},
		{
			Name:        "T3",
			Environment: "stress-test", // fail
			CpuCores:    0,             // fail
			Ram:         "255",         // fail
		},
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	for _, input := range inputs {
		terraformOptions.Vars = map[string]interface{}{
			"name":        input.Name,
			"environment": input.Environment,
			"cpu_cores":   input.CpuCores,
			"ram":         input.Ram,
		}

		_, err := terraform.InitAndApplyE(t, terraformOptions)

		assert.Error(t, err)
	}
}
