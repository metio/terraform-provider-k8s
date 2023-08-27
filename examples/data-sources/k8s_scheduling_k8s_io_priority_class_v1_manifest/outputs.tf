output "manifests" {
  value = {
    "example" = data.k8s_scheduling_k8s_io_priority_class_v1_manifest.example.yaml
  }
}
