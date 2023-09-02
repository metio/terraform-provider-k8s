output "manifests" {
  value = {
    "example" = data.k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1_manifest.example.yaml
  }
}
