data "k8s_opentelemetry_io_op_amp_bridge_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
