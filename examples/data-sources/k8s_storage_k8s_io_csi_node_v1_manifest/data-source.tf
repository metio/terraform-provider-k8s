data "k8s_storage_k8s_io_csi_node_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    drivers = []
  }
}
