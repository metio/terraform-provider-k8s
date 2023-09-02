data "k8s_ceph_rook_io_ceph_filesystem_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
