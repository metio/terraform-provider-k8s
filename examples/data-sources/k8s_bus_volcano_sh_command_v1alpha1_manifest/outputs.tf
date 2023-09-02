output "manifests" {
  value = {
    "example" = data.k8s_bus_volcano_sh_command_v1alpha1_manifest.example.yaml
  }
}
