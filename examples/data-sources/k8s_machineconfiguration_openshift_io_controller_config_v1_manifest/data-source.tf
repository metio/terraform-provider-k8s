data "k8s_machineconfiguration_openshift_io_controller_config_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
