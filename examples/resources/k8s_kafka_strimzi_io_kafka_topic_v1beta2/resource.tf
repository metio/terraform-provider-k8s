resource "k8s_kafka_strimzi_io_kafka_topic_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
}
