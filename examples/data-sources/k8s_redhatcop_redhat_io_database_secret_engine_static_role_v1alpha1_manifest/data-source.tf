data "k8s_redhatcop_redhat_io_database_secret_engine_static_role_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
