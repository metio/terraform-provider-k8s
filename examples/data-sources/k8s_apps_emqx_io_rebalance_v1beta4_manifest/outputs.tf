output "manifests" {
  value = {
    "example" = data.k8s_apps_emqx_io_rebalance_v1beta4_manifest.example.yaml
  }
}
