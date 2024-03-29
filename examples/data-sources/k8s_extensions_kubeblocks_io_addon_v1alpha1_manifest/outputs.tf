output "manifests" {
  value = {
    "example" = data.k8s_extensions_kubeblocks_io_addon_v1alpha1_manifest.example.yaml
  }
}
