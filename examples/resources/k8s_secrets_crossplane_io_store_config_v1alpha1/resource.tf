resource "k8s_secrets_crossplane_io_store_config_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    default_scope = "local"
  }
}
