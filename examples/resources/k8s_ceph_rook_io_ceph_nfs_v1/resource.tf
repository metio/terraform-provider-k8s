resource "k8s_ceph_rook_io_ceph_nfs_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
