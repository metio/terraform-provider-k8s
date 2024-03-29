data "k8s_capabilities_3scale_net_active_doc_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
