output "manifests" {
  value = {
    "example" = data.k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1_manifest.example.yaml
  }
}
