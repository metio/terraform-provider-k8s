data "k8s_cilium_io_cilium_identity_v2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  security_labels = {}
}
