resource "k8s_ceph_rook_io_ceph_object_zone_group_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    realm = "zone-a"
  }
}
