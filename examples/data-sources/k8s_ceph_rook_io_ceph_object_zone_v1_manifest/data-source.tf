data "k8s_ceph_rook_io_ceph_object_zone_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    data_pool     = {}
    metadata_pool = {}
    zone_group    = "group-1"
  }
}
