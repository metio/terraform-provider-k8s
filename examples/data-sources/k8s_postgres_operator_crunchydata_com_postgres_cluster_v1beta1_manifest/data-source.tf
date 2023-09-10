data "k8s_postgres_operator_crunchydata_com_postgres_cluster_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    postgres_version = 12
    instances        = []
    backups = {
      pgbackrest = {
        repos = [
          {
            name = "repo1"
          }
        ]
      }
    }
  }
}
