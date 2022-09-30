resource "k8s_sparkoperator_k8s_io_scheduled_spark_application_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    schedule = "some-schedule"
    template = {
      driver        = {}
      executor      = {}
      spark_version = "some-version"
      type          = "some-type"
    }
  }
}
