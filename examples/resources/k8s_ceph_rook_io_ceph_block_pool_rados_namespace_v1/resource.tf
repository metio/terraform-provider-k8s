resource "k8s_ceph_rook_io_ceph_block_pool_rados_namespace_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    block_pool_name = "pool-b"
  }
}
