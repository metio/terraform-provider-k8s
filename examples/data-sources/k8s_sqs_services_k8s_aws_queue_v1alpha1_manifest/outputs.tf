output "manifests" {
  value = {
    "example" = data.k8s_sqs_services_k8s_aws_queue_v1alpha1_manifest.example.yaml
  }
}
