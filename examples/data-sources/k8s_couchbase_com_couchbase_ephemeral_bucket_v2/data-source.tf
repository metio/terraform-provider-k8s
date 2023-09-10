data "k8s_couchbase_com_couchbase_ephemeral_bucket_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
