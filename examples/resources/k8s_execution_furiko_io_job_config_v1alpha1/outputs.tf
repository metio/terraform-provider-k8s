output "resources" {
  value = {
    "minimal" = k8s_execution_furiko_io_job_config_v1alpha1.minimal.yaml
    "example" = k8s_execution_furiko_io_job_config_v1alpha1.example.yaml
  }
}
