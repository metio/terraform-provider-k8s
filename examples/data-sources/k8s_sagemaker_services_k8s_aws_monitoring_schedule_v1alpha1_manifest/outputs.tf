output "manifests" {
  value = {
    "example" = data.k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest.example.yaml
  }
}
