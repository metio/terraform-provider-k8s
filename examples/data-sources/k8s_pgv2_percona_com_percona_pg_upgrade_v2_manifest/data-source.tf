data "k8s_pgv2_percona_com_percona_pg_upgrade_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    from_postgres_version = 15
    to_postgres_version   = 16
    postgres_cluster_name = "some-cluster"
    image                 = "some-image"
  }
}
