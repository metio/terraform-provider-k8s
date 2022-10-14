resource "k8s_apm_k8s_elastic_co_apm_server_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_apm_k8s_elastic_co_apm_server_v1" "example" {
  metadata = {
    name = "apmserver-sample"
  }
  spec = {
    version = "8.4.0"
    count   = 1
    elasticsearch_ref = {
      name = "elasticsearch-sample"
    }
  }
}
