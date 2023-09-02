data "k8s_sparkoperator_k8s_io_spark_application_v1beta2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
