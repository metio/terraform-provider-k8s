output "manifests" {
  value = {
    "example" = data.k8s_policy_karmada_io_federated_resource_quota_v1alpha1_manifest.example.yaml
  }
}
