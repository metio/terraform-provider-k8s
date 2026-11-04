data "k8s_api_clever_cloud_com_my_sql_v1_manifest" "example" {
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
      version    = 57
    }
    organisation = "some-value"
  }
}
