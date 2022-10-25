resource "k8s_ceph_rook_io_ceph_client_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    caps = 3
  }
}
