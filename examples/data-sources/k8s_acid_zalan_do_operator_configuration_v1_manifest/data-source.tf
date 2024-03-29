data "k8s_acid_zalan_do_operator_configuration_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  configuration = {
    max_instances = 7
    min_instances = 3
  }
}
