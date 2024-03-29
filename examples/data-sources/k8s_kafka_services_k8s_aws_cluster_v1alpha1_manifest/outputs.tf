output "manifests" {
  value = {
    "example" = data.k8s_kafka_services_k8s_aws_cluster_v1alpha1_manifest.example.yaml
  }
}
