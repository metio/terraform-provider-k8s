data "k8s_apm_k8s_elastic_co_apm_server_v1_manifest" "example" {
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
