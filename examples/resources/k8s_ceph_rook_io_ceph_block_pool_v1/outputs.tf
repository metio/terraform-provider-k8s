output "resources" {
  value = {
    "minimal" = k8s_ceph_rook_io_ceph_block_pool_v1.minimal.yaml
  }
}
