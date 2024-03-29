output "manifests" {
  value = {
    "example" = data.k8s_apps_emqx_io_emqx_broker_v1beta4_manifest.example.yaml
  }
}
