resource "k8s_ceph_rook_io_ceph_filesystem_sub_volume_group_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    filesystem_name = "hank"
  }
}
