output "resources" {
  value = {
    "minimal" = k8s_argoproj_io_application_v1alpha1.minimal.yaml
  }
}
