data "k8s_ceph_rook_io_ceph_bucket_topic_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
