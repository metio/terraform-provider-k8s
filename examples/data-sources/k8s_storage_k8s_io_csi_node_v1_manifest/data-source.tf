data "k8s_storage_k8s_io_csi_node_v1_manifest" "example" {
  metadata = {
    name = "some-name"
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
