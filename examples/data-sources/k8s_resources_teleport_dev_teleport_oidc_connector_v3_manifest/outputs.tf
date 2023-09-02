output "manifests" {
  value = {
    "example" = data.k8s_resources_teleport_dev_teleport_oidc_connector_v3_manifest.example.yaml
  }
}
