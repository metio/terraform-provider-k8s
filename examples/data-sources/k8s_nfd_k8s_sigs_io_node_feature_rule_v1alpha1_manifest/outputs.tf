output "manifests" {
  value = {
    "example" = data.k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest.example.yaml
  }
}
