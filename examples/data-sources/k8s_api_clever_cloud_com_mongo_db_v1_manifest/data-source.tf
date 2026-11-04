data "k8s_api_clever_cloud_com_mongo_db_v1_manifest" "example" {
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
      encryption = false
      version    = 403
    }
    organisation = "some-value"
  }
}
