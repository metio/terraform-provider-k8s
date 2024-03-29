output "manifests" {
  value = {
    "example" = data.k8s_kamaji_clastix_io_data_store_v1alpha1_manifest.example.yaml
  }
}
