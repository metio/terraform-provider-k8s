output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_kubernetes_endpoint_resolver_v2_manifest.example.yaml
  }
}
