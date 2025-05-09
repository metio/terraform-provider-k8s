output "manifests" {
  value = {
    "example" = data.k8s_amd_com_device_config_v1alpha1_manifest.example.yaml
  }
}
