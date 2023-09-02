output "manifests" {
  value = {
    "example" = data.k8s_kafka_strimzi_io_kafka_topic_v1alpha1_manifest.example.yaml
  }
}
