resource "k8s_couchbase_com_couchbase_autoscaler_v2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}