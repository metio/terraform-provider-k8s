resource "k8s_mariadb_mmontes_io_database_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
