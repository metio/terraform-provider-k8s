output "manifests" {
  value = {
    "example" = data.k8s_extensions_istio_io_wasm_plugin_v1alpha1_manifest.example.yaml
  }
}
