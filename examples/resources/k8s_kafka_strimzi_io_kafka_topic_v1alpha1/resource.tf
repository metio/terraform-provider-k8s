resource "k8s_kafka_strimzi_io_kafka_topic_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
