data "k8s_mariadb_persistentsys_maria_db_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
