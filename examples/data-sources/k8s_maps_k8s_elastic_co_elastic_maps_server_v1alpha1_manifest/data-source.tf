data "k8s_maps_k8s_elastic_co_elastic_maps_server_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    version = "8.4.0"
    count   = 1
    elasticsearch_ref = {
      name = "elasticsearch-sample"
    }
  }
}
