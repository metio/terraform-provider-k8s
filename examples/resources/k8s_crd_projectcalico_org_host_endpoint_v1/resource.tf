resource "k8s_crd_projectcalico_org_host_endpoint_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
