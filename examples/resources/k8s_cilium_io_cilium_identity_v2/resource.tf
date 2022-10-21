resource "k8s_cilium_io_cilium_identity_v2" "minimal" {
  metadata = {
    name = "test"
  }
  security__labels = {}
}
