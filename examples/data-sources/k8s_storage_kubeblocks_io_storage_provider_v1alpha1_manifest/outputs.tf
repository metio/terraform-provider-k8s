output "manifests" {
  value = {
    "example" = data.k8s_storage_kubeblocks_io_storage_provider_v1alpha1_manifest.example.yaml
  }
}
