data "k8s_redhatcop_redhat_io_vault_secret_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
