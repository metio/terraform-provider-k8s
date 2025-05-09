output "manifests" {
  value = {
    "example" = data.k8s_akri_sh_instance_v0_manifest.example.yaml
  }
}
