output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_kubernetes_service_resolver_v2_manifest.example.yaml
  }
}
