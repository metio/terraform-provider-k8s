output "manifests" {
  value = {
    "example" = data.k8s_apps_emqx_io_rebalance_v2beta1_manifest.example.yaml
  }
}
