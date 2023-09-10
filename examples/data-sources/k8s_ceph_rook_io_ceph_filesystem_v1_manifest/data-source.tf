data "k8s_ceph_rook_io_ceph_filesystem_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    data_pools    = []
    metadata_pool = {}
    metadata_server = {
      active_count = 2
    }
  }
}
