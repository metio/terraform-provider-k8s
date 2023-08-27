data "k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1" "example" {
  metadata = {
    name = "some-name"

  }
}
