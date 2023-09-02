output "manifests" {
  value = {
    "example" = data.k8s_maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest.example.yaml
  }
}
