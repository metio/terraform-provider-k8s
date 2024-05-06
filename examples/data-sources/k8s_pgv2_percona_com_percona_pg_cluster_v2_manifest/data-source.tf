data "k8s_pgv2_percona_com_percona_pg_cluster_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    backups = {
      pgbackrest = {
        repos = []
      }
    }
    instances        = []
    postgres_version = "15"
  }
}
