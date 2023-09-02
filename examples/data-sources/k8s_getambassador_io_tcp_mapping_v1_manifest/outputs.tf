output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_tcp_mapping_v1_manifest.example.yaml
  }
}
