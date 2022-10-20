output "resources" {
  value = {
    "minimal"         = k8s_execution_furiko_io_job_v1alpha1.minimal.yaml
    "from_job_config" = k8s_execution_furiko_io_job_v1alpha1.from_job_config.yaml
    "standalone"      = k8s_execution_furiko_io_job_v1alpha1.standalone.yaml
  }
}
