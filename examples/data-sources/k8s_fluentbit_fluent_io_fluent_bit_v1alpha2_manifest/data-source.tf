data "k8s_fluentbit_fluent_io_fluent_bit_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
