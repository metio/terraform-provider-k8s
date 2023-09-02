data "k8s_ceph_rook_io_ceph_cosi_driver_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
