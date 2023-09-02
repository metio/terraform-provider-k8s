output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_dev_portal_v2_manifest.example.yaml
  }
}
