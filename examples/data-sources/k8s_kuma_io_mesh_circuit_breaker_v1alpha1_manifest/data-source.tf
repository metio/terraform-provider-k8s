data "k8s_kuma_io_mesh_circuit_breaker_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
