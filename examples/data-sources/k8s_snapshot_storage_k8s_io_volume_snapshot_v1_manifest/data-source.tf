data "k8s_snapshot_storage_k8s_io_volume_snapshot_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    source = {}
  }
}
