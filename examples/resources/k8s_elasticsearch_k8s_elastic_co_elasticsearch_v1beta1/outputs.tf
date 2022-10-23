output "resources" {
  value = {
    "minimal" = k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1beta1.minimal.yaml
    "example" = k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1beta1.example.yaml
  }
}
