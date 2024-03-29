output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_admission_report_v2_manifest.example.yaml
  }
}
