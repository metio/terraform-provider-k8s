resource "k8s_external_secrets_io_cluster_secret_store_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
