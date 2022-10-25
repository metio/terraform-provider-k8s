resource "k8s_ceph_rook_io_ceph_bucket_topic_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    endpoint               = {}
    object_store_name      = "store-x"
    object_store_namespace = "storage"
  }
}
