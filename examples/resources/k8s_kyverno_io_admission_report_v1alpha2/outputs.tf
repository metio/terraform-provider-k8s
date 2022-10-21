output "resources" {
  value = {
    "minimal" = k8s_kyverno_io_admission_report_v1alpha2.minimal.yaml
  }
}
