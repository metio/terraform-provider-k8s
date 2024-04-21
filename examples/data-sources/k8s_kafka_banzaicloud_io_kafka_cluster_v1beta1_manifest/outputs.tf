output "manifests" {
  value = {
    "example" = data.k8s_kafka_banzaicloud_io_kafka_cluster_v1beta1_manifest.example.yaml
  }
}
