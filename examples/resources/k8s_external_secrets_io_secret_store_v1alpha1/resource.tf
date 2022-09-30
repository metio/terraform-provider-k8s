resource "k8s_external_secrets_io_secret_store_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
