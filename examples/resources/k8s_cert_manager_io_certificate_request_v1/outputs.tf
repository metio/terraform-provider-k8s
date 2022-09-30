output "resources" {
  value = {
    "minimal" = k8s_cert_manager_io_certificate_request_v1.minimal.yaml
  }
}
