output "resources" {
  value = {
    "minimal" = k8s_operator_tigera_io_image_set_v1.minimal.yaml
  }
}
