output "manifests" {
  value = {
    "example" = data.k8s_apps_stateful_set_v1_manifest.example.yaml
  }
}
