output "manifests" {
  value = {
    "example" = data.k8s_operator_cluster_x_k8s_io_control_plane_provider_v1alpha2_manifest.example.yaml
  }
}
