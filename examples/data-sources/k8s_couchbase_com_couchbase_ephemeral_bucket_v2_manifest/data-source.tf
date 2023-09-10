data "k8s_couchbase_com_couchbase_ephemeral_bucket_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
