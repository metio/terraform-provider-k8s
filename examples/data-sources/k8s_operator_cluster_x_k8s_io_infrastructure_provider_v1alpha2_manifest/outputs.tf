output "manifests" {
  value = {
    "example" = data.k8s_operator_cluster_x_k8s_io_infrastructure_provider_v1alpha2_manifest.example.yaml
  }
}
