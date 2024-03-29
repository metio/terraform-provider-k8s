data "k8s_kibana_k8s_elastic_co_kibana_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
