resource "k8s_ceph_rook_io_ceph_nfs_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    server = {
      active = 3
    }
  }
}
