resource "k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1" "example" {
  metadata = {
    name = "ent-sample"
  }
  spec = {
    version = "8.4.0"
    config = {
      "ent_search.external_url" = "https://localhost:3002"
    }
    count = 1
    elasticsearch_ref = {
      name = "elasticsearch-sample"
    }
  }
}
