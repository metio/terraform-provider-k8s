data "k8s_objectstorage_k8s_io_bucket_access_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
