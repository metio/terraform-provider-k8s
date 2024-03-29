data "k8s_operator_openshift_io_ingress_controller_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
