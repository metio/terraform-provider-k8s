data "k8s_objectbucket_io_object_bucket_claim_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
