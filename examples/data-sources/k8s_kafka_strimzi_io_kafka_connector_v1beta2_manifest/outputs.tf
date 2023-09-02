output "manifests" {
  value = {
    "example" = data.k8s_kafka_strimzi_io_kafka_connector_v1beta2_manifest.example.yaml
  }
}
