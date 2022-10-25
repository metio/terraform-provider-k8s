resource "k8s_ceph_rook_io_ceph_object_realm_v1" "minimal" {
  metadata = {
    name = "test"
  }
}
