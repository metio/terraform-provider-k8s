output "resources" {
  value = {
    "minimal" = k8s_networking_istio_io_virtual_service_v1beta1.minimal.yaml
  }
}
