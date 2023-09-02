output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_tcp_mapping_v3alpha1_manifest.example.yaml
  }
}
