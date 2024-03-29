output "manifests" {
  value = {
    "example" = data.k8s_cluster_ipfs_io_ipfs_cluster_v1alpha1_manifest.example.yaml
  }
}
