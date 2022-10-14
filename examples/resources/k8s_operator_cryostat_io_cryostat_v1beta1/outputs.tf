output "resources" {
  value = {
    "minimal" = k8s_operator_cryostat_io_cryostat_v1beta1.minimal.yaml
    "example" = k8s_operator_cryostat_io_cryostat_v1beta1.example.yaml
  }
}
