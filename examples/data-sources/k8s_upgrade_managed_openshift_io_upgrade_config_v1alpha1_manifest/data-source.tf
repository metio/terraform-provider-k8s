data "k8s_upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
