data "k8s_postgresql_cnpg_io_pooler_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster = {
      name = "some-name"
    }
    pgbouncer = {}
  }
}
