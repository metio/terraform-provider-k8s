output "manifests" {
  value = {
    "example" = data.k8s_stunner_l7mp_io_gateway_config_v1_manifest.example.yaml
  }
}
