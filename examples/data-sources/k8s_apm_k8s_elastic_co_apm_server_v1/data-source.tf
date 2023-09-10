data "k8s_apm_k8s_elastic_co_apm_server_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
