output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_module_v1_manifest.example.yaml
  }
}
