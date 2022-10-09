resource "k8s_sparkoperator_k8s_io_spark_application_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    driver        = {}
    executor      = {}
    spark_version = "some-version"
    type          = "Java"
  }
}
