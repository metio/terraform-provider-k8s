output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_tls_context_v2_manifest.example.yaml
  }
}
