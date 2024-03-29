output "manifests" {
  value = {
    "example" = data.k8s_cluster_ipfs_io_circuit_relay_v1alpha1_manifest.example.yaml
  }
}
