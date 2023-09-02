output "manifests" {
  value = {
    "example" = data.k8s_gloo_solo_io_upstream_group_v1_manifest.example.yaml
  }
}
