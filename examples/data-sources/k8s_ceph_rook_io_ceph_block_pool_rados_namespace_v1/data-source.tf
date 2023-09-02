data "k8s_ceph_rook_io_ceph_block_pool_rados_namespace_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
