output "manifests" {
  value = {
    "example" = data.k8s_marin3r_3scale_net_envoy_config_v1alpha1_manifest.example.yaml
  }
}
