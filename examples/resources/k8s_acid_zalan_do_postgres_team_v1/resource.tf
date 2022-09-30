resource "k8s_acid_zalan_do_postgres_team_v1" "big" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
    labels = {
      "test" = "abc"
    }
    annotations = {
      "try" = "this"
    }
  }
  spec = {
    additional_members = {
      "team-a" = ["bob", "bill", "barry"]
    }
    additional_superuser_teams = {
      "team-b" = ["alice", "eve", "julia"]
    }
    additional_teams = {
      "team-c" = ["team-1", "team-2", "team-3"]
    }
  }
}

resource "k8s_acid_zalan_do_postgres_team_v1" "small" {
  metadata = {
    name = "test"
  }
  spec = {
    additional_members = {
      "team-a" = ["bob", "bill", "barry"]
    }
  }
}
