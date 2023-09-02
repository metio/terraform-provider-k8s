output "manifests" {
  value = {
    "example" = data.k8s_resources_teleport_dev_teleport_user_v2_manifest.example.yaml
  }
}
