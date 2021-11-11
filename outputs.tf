output "metadata" {
  value = {
    server_name = local.server_name
    server_uid  = local.server_uid
    spec        = local.spec
  }
  description = "Server metadata."
}