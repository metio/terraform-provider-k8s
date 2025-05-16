output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_ldap_auth_engine_config_v1alpha1_manifest.example.yaml
  }
}
