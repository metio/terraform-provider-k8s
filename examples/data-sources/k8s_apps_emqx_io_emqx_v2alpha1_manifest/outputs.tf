output "manifests" {
  value = {
    "example" = data.k8s_apps_emqx_io_emqx_v2alpha1_manifest.example.yaml
  }
}
