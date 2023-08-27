data "k8s_storage_k8s_io_volume_attachment_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    attacher  = "some-name"
    node_name = "some-node"
    source = {
      persistent_volume_name = "some-volume"
    }
  }
}
