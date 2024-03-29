data "k8s_azure_microsoft_com_key_vault_key_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
