data "k8s_ingress_operator_openshift_io_dns_record_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
