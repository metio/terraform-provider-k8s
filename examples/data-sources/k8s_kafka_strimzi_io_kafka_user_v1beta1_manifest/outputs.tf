output "manifests" {
  value = {
    "example" = data.k8s_kafka_strimzi_io_kafka_user_v1beta1_manifest.example.yaml
  }
}
