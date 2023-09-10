data "k8s_databases_schemahero_io_database_v1alpha4_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {}
}
