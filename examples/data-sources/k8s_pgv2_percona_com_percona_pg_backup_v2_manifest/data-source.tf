data "k8s_pgv2_percona_com_percona_pg_backup_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    pg_cluster = "some-cluster"
    repo_name  = "repo1"
  }
}
