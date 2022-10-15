output "resources" {
  value = {
    "minimal" = k8s_batch_cron_job_v1.minimal.yaml
    "example" = k8s_batch_cron_job_v1.example.yaml
  }
}
