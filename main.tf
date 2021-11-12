# here just to mimic data processing and some resource outputs.

resource "random_string" "server_uid" {
  length           = 16
  special          = true
  override_special = "/@£$"
}

resource "random_string" "generated_server_name" {
  count = length(var.name) > 0 ? 0 : 1

  length           = 8
  special          = true
  override_special = "/@£$"
}

locals {
  server_name = length(var.name) > 0 ? upper("${var.name}-${var.environment}") : upper("${random_string.generated_server_name[0].result}-${var.environment}")
  server_uid  = upper(random_string.server_uid.result)
  spec = {
    cpu_cores = var.cpu_cores
    ram       = var.ram
  }
}