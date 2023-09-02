output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_recurring_job_v1beta1_manifest.example.yaml
  }
}
