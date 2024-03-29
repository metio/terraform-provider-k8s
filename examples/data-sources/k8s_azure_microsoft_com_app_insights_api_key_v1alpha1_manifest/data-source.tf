data "k8s_azure_microsoft_com_app_insights_api_key_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
