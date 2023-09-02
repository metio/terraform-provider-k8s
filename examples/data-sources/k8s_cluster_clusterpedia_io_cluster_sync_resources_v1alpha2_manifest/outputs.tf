output "manifests" {
  value = {
    "example" = data.k8s_cluster_clusterpedia_io_cluster_sync_resources_v1alpha2_manifest.example.yaml
  }
}
