output "manifests" {
  value = {
    "example" = data.k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest.example.yaml
  }
}
