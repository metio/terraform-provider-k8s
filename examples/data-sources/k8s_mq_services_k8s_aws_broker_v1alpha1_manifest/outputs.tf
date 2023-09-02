output "manifests" {
  value = {
    "example" = data.k8s_mq_services_k8s_aws_broker_v1alpha1_manifest.example.yaml
  }
}
