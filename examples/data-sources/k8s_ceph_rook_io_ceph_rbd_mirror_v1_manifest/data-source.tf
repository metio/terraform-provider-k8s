data "k8s_ceph_rook_io_ceph_rbd_mirror_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    count = 7
  }
}
