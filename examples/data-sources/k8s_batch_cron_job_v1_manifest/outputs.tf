output "manifests" {
  value = {
    "example" = data.k8s_batch_cron_job_v1_manifest.example.yaml
  }
}
