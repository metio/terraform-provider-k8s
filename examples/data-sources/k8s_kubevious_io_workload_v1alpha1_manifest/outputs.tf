output "manifests" {
  value = {
    "example" = data.k8s_kubevious_io_workload_v1alpha1_manifest.example.yaml
  }
}
