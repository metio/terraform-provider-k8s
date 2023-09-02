data "k8s_ceph_rook_io_ceph_object_realm_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
