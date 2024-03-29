data "k8s_azure_microsoft_com_azure_public_ip_address_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
