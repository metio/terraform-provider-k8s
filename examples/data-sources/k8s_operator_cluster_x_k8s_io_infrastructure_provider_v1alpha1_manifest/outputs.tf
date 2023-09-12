output "manifests" {
  value = {
    "example" = data.k8s_operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1_manifest.example.yaml
  }
}
