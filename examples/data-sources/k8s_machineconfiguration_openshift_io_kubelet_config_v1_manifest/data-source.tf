data "k8s_machineconfiguration_openshift_io_kubelet_config_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
