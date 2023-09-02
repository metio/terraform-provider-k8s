data "k8s_sparkoperator_k8s_io_scheduled_spark_application_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
