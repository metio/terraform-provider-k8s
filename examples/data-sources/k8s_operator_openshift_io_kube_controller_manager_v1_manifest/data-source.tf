data "k8s_operator_openshift_io_kube_controller_manager_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
