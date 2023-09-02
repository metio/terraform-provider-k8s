resource "k8s_couchbase_com_couchbase_memcached_bucket_v2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
