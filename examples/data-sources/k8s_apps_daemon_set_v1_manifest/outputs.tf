output "manifests" {
  value = {
    "example" = data.k8s_apps_daemon_set_v1_manifest.example.yaml
  }
}
