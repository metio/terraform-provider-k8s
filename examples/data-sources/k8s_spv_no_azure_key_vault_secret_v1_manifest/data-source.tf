data "k8s_spv_no_azure_key_vault_secret_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
