data "k8s_operator_openshift_io_kube_api_server_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
