output "resources" {
  value = {
    "minimal" = k8s_apm_k8s_elastic_co_apm_server_v1beta1.minimal.yaml
    "example" = k8s_apm_k8s_elastic_co_apm_server_v1beta1.example.yaml
  }
}
