output "manifests" {
  value = {
    "example" = data.k8s_wgpolicyk8s_io_policy_report_v1alpha1_manifest.example.yaml
  }
}
