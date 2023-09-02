output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_module_v2_manifest.example.yaml
  }
}
