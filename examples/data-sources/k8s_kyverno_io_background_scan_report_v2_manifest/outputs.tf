output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_background_scan_report_v2_manifest.example.yaml
  }
}
