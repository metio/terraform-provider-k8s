resource "k8s_ceph_rook_io_ceph_filesystem_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    data_pools    = []
    metadata_pool = {}
    metadata_server = {
      active_count = 2
    }
  }
}
