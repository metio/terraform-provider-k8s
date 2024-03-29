data "k8s_azure_microsoft_com_api_mgmt_api_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
