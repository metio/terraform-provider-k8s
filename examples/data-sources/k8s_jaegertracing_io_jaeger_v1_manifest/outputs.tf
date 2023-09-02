output "manifests" {
  value = {
    "example" = data.k8s_jaegertracing_io_jaeger_v1_manifest.example.yaml
  }
}
