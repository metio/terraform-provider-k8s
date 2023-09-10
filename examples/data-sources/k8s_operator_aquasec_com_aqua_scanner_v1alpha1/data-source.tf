data "k8s_operator_aquasec_com_aqua_scanner_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
