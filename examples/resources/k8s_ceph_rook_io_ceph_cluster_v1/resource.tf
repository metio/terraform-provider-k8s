resource "k8s_ceph_rook_io_ceph_cluster_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
