data "k8s_elasticsearch_k8s_elastic_co_elasticsearch_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
