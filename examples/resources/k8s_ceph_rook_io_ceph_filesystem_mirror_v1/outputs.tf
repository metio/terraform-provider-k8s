output "resources" {
  value = {
    "minimal" = k8s_ceph_rook_io_ceph_filesystem_mirror_v1.minimal.yaml
  }
}
