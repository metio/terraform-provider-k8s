output "manifests" {
  value = {
    "example" = data.k8s_networking_istio_io_workload_entry_v1_manifest.example.yaml
  }
}
