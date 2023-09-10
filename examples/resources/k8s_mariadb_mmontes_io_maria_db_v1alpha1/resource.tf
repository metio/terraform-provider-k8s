resource "k8s_mariadb_mmontes_io_maria_db_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
