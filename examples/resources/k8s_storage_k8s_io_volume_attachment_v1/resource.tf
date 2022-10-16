resource "k8s_storage_k8s_io_volume_attachment_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    attacher  = "some-name"
    node_name = "some-node"
    source = {
      persistent_volume_name = "some-volume"
    }
  }
}
