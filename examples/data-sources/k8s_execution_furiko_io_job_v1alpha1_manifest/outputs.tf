output "manifests" {
  value = {
    "example" = data.k8s_execution_furiko_io_job_v1alpha1_manifest.example.yaml
  }
}
