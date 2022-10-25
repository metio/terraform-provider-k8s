resource "k8s_ceph_rook_io_ceph_object_zone_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    data_pool     = {}
    metadata_pool = {}
    zone_group    = "group-1"
  }
}
