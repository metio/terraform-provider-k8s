output "manifests" {
  value = {
    "example" = data.k8s_enterprise_gloo_solo_io_auth_config_v1_manifest.example.yaml
  }
}
