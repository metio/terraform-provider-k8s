output "resources" {
  value = {
    "minimal" = k8s_kafka_strimzi_io_kafka_user_v1beta1.minimal.yaml
  }
}
