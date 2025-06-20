data "k8s_logstash_k8s_elastic_co_logstash_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
