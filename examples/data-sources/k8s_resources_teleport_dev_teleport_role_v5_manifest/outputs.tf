output "manifests" {
  value = {
    "example" = data.k8s_resources_teleport_dev_teleport_role_v5_manifest.example.yaml
  }
}
