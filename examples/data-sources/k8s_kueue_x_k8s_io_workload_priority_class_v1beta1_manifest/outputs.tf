output "manifests" {
  value = {
    "example" = data.k8s_kueue_x_k8s_io_workload_priority_class_v1beta1_manifest.example.yaml
  }
}
