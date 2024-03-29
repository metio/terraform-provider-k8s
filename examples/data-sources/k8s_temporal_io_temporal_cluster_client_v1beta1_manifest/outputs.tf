output "manifests" {
  value = {
    "example" = data.k8s_temporal_io_temporal_cluster_client_v1beta1_manifest.example.yaml
  }
}
