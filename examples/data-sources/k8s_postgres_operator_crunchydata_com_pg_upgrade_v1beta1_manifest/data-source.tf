data "k8s_postgres_operator_crunchydata_com_pg_upgrade_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
