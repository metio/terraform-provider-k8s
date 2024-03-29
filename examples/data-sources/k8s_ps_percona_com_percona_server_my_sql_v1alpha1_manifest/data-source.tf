data "k8s_ps_percona_com_percona_server_my_sql_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
