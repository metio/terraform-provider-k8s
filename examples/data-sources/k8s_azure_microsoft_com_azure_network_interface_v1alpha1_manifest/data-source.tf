data "k8s_azure_microsoft_com_azure_network_interface_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
