output "manifests" {
  value = {
    "example" = data.k8s_wgpolicyk8s_io_cluster_policy_report_v1beta1_manifest.example.yaml
  }
}
