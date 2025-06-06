output "manifests" {
  value = {
    "example" = data.k8s_logstash_k8s_elastic_co_logstash_v1alpha1_manifest.example.yaml
  }
}
