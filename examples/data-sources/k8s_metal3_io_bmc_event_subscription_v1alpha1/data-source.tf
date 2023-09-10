data "k8s_metal3_io_bmc_event_subscription_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
