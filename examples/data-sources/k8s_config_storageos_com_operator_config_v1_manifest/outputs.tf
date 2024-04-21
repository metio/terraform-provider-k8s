output "manifests" {
  value = {
    "example" = data.k8s_config_storageos_com_operator_config_v1_manifest.example.yaml
  }
}
