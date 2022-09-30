output "resources" {
  value = {
    "minimal" = k8s_jaegertracing_io_jaeger_v1.minimal.yaml
  }
}
