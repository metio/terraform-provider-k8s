output "manifests" {
  value = {
    "example" = data.k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1_manifest.example.yaml
  }
}
