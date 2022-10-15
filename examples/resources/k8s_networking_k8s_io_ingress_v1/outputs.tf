output "resources" {
  value = {
    "minimal" = k8s_networking_k8s_io_ingress_v1.minimal.yaml
    "example" = k8s_networking_k8s_io_ingress_v1.example.yaml
  }
}
