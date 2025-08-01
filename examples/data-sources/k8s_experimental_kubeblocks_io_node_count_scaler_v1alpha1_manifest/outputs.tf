output "manifests" {
  value = {
    "example" = data.k8s_experimental_kubeblocks_io_node_count_scaler_v1alpha1_manifest.example.yaml
  }
}
