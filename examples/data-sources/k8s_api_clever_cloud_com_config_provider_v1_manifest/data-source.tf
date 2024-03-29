data "k8s_api_clever_cloud_com_config_provider_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
