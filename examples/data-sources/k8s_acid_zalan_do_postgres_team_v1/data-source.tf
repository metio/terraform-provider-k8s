data "k8s_acid_zalan_do_postgres_team_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
