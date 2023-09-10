data "k8s_ceph_rook_io_ceph_bucket_topic_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    endpoint               = {}
    object_store_name      = "store-x"
    object_store_namespace = "storage"
  }
}
