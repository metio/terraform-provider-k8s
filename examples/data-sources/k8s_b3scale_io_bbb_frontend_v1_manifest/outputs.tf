output "manifests" {
  value = {
    "example" = data.k8s_b3scale_io_bbb_frontend_v1_manifest.example.yaml
  }
}
