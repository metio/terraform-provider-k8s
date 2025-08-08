output "manifests" {
  value = {
    "example" = data.k8s_reports_x_k8s_io_cluster_policy_report_v1beta2_manifest.example.yaml
  }
}
