data "k8s_acid_zalan_do_postgres_team_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    additional_members = {
      "team-a" = ["bob", "bill", "barry"]
    }
  }
}
