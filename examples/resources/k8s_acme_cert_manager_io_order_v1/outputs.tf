output "resources" {
  value = {
    "minimal" = k8s_acme_cert_manager_io_order_v1.minimal.yaml
  }
}
