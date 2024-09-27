output "manifests" {
  value = {
    "example" = data.k8s_kueue_x_k8s_io_cohort_v1alpha1_manifest.example.yaml
  }
}
