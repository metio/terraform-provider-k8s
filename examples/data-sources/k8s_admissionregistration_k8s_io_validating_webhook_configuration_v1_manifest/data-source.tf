data "k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
