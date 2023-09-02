output "manifests" {
  value = {
    "example" = data.k8s_work_karmada_io_work_v1alpha1_manifest.example.yaml
  }
}
