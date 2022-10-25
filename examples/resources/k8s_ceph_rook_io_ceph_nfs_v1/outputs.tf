output "resources" {
  value = {
    "minimal" = k8s_ceph_rook_io_ceph_nfs_v1.minimal.yaml
  }
}
