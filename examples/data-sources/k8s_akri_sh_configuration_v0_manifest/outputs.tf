output "manifests" {
  value = {
    "example" = data.k8s_akri_sh_configuration_v0_manifest.example.yaml
  }
}
