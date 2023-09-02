output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_l2_announcement_policy_v2alpha1_manifest.example.yaml
  }
}
