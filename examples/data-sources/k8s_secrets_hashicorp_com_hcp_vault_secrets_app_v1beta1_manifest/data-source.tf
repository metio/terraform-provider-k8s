data "k8s_secrets_hashicorp_com_hcp_vault_secrets_app_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
