output "manifests" {
  value = {
    "example" = data.k8s_logging_extensions_banzaicloud_io_event_tailer_v1alpha1_manifest.example.yaml
  }
}
