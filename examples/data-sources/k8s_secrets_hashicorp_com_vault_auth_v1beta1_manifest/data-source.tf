data "k8s_secrets_hashicorp_com_vault_auth_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
