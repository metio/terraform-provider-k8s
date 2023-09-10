data "k8s_schemas_schemahero_io_migration_v1alpha4_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
