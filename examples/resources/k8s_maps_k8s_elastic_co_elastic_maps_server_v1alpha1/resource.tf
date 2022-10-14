resource "k8s_maps_k8s_elastic_co_elastic_maps_server_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_maps_k8s_elastic_co_elastic_maps_server_v1alpha1" "example" {
  metadata = {
    name = "ems-sample"
  }
  spec = {
    version = "8.4.0"
    count   = 1
    elasticsearch_ref = {
      name = "elasticsearch-sample"
    }
  }
}
