output "manifests" {
  value = {
    "example" = data.k8s_appprotectdos_f5_com_dos_protected_resource_v1beta1_manifest.example.yaml
  }
}
