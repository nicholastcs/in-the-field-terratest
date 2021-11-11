package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// by default, this should pass, duh.
func TestTerraformValidate(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	terraform.Validate(t, terraformOptions)
}

func TestModuleSensibleDefault(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOptions)

	output := struct {
		ServerName string `json:"server_name"`
		ServerUid  string `json:"server_uid"`
		Spec       struct {
			CpuCores int    `json:"cpu_cores"`
			Ram      string `json:"ram"`
		} `json:"spec"`
	}{}

	terraform.InitAndApply(t, terraformOptions)
	terraform.OutputStruct(t, terraformOptions, "metadata", &output)

	want := struct {
		CpuCores int    `json:"cpu_cores"`
		Ram      string `json:"ram"`
	}{
		CpuCores: 2,
		Ram:      "256",
	}

	got := output.Spec

	assert.Equal(t, want, got)
	assert.NotEmpty(t, output.ServerUid)
	assert.Equal(t, true, strings.HasSuffix(output.ServerName, "-DEV"))
}
