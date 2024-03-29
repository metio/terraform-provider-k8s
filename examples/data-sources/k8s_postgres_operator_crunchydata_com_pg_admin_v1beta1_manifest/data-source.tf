data "k8s_postgres_operator_crunchydata_com_pg_admin_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
