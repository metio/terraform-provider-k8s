output "manifests" {
  value = {
    "example" = data.k8s_slo_koordinator_sh_node_slo_v1alpha1_manifest.example.yaml
  }
}
