output "manifests" {
  value = {
    "example" = data.k8s_resources_teleport_dev_teleport_github_connector_v3_manifest.example.yaml
  }
}
