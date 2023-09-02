output "manifests" {
  value = {
    "example" = data.k8s_agent_k8s_elastic_co_agent_v1alpha1_manifest.example.yaml
  }
}
