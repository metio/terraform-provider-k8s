output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_cluster_admission_report_v1alpha2_manifest.example.yaml
  }
}
