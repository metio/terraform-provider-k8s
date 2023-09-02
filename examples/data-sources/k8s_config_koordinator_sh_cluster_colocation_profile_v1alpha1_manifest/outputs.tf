output "manifests" {
  value = {
    "example" = data.k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1_manifest.example.yaml
  }
}
