data "k8s_snapshot_storage_k8s_io_volume_snapshot_content_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    deletion_policy     = "Retain"
    driver              = "some-driver"
    source              = {}
    volume_snapshot_ref = {}
  }
}
