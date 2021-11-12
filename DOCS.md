# Terraform Module Documentation

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=0.15.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | 3.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_random"></a> [random](#provider\_random) | 3.1.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [random_string.generated_server_name](https://registry.terraform.io/providers/hashicorp/random/3.1.0/docs/resources/string) | resource |
| [random_string.server_uid](https://registry.terraform.io/providers/hashicorp/random/3.1.0/docs/resources/string) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cpu_cores"></a> [cpu\_cores](#input\_cpu\_cores) | Number of CPU cores. | `number` | `2` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | Environment tag. | `string` | `"dev"` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of VM. Random string will generated if it is empty. | `string` | `""` | no |
| <a name="input_ram"></a> [ram](#input\_ram) | Number of CPU cores. | `string` | `"256"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_metadata"></a> [metadata](#output\_metadata) | Server metadata. |
<!-- END_TF_DOCS -->