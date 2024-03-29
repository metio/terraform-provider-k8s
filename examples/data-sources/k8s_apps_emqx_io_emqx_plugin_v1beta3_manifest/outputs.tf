output "manifests" {
  value = {
    "example" = data.k8s_apps_emqx_io_emqx_plugin_v1beta3_manifest.example.yaml
  }
}
