resource "k8s_kafka_strimzi_io_kafka_topic_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
