output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_mapping_v2_manifest.example.yaml
  }
}
