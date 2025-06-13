data "k8s_generators_external_secrets_io_generator_state_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
