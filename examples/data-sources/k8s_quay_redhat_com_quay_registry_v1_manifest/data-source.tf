data "k8s_quay_redhat_com_quay_registry_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
