data "k8s_pxc_percona_com_percona_xtra_db_cluster_restore_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
