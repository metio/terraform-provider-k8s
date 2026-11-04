data "k8s_api_clever_cloud_com_elastic_search_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  spec = {
    instance = {
      plan   = "some-value"
      region = "some-value"
    }
    options = {
      apm        = false
      encryption = false
      kibana     = false
      version    = 6
    }
    organisation = "some-value"
  }
}
