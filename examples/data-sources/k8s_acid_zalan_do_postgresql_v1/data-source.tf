data "k8s_acid_zalan_do_postgresql_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
