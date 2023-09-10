data "k8s_acid_zalan_do_postgresql_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    number_of_instances = 3
    postgresql = {
      version = "9.6"
    }
    team_id = "abc"
    volume = {
      storage_class = "gp3"
      selector = {
        match_labels = {
          "app.kubernetes.io/name" = "some-example"
        }
      }
      size = "17G"
    }
  }
}
