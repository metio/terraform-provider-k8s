data "k8s_k8s_otterize_com_client_intents_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
