output "resources" {
  value = {
    "minimal" = k8s_kibana_k8s_elastic_co_kibana_v1.minimal.yaml
    "example" = k8s_kibana_k8s_elastic_co_kibana_v1.example.yaml
  }
}
