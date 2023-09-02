data "k8s_kafka_strimzi_io_kafka_user_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
