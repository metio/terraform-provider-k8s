output "manifests" {
  value = {
    "example" = data.k8s_multicluster_x_k8s_io_applied_work_v1alpha1_manifest.example.yaml
  }
}
