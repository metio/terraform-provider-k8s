data "k8s_jaegertracing_io_jaeger_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
