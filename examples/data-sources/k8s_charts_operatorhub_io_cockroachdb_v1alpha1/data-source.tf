data "k8s_charts_operatorhub_io_cockroachdb_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
