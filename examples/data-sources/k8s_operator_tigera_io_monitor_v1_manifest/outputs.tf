output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_monitor_v1_manifest.example.yaml
  }
}
