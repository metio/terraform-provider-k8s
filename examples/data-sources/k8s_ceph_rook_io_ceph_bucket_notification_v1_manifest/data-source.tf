data "k8s_ceph_rook_io_ceph_bucket_notification_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    topic = "hotnews"
  }
}
