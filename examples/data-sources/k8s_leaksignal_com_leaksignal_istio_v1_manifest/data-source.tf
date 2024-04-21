data "k8s_leaksignal_com_leaksignal_istio_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
