data "k8s_projectcontour_io_tls_certificate_delegation_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
