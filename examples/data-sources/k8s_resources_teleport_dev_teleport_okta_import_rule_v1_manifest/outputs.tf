output "manifests" {
  value = {
    "example" = data.k8s_resources_teleport_dev_teleport_okta_import_rule_v1_manifest.example.yaml
  }
}
