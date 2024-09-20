output "manifests" {
  value = {
    "example" = data.k8s_kafka_stackable_tech_kafka_cluster_v1alpha1_manifest.example.yaml
  }
}
