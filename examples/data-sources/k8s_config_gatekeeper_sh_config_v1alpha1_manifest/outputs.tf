output "manifests" {
  value = {
    "example" = data.k8s_config_gatekeeper_sh_config_v1alpha1_manifest.example.yaml
  }
}
