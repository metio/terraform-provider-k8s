output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_cert_auth_engine_config_v1alpha1_manifest.example.yaml
  }
}
