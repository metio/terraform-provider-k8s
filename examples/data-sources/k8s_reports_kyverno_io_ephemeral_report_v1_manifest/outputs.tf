output "manifests" {
  value = {
    "example" = data.k8s_reports_kyverno_io_ephemeral_report_v1_manifest.example.yaml
  }
}
