output "resources" {
  value = {
    "minimal" = k8s_kyverno_io_generate_request_v1.minimal.yaml
  }
}
