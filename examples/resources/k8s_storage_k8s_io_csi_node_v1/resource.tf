resource "k8s_storage_k8s_io_csi_node_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    drivers = []
  }
}

resource "k8s_storage_k8s_io_csi_node_v1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    drivers = [
      {
        name         = "some-name"
        node_id      = "some-node"
        topologyKeys = ["io.kubernetes/zone"]
      }
    ]
  }
}
