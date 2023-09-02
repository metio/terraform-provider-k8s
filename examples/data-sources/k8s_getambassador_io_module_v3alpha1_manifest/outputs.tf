output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_module_v3alpha1_manifest.example.yaml
  }
}
