data "k8s_ocmagent_managed_openshift_io_managed_fleet_notification_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
