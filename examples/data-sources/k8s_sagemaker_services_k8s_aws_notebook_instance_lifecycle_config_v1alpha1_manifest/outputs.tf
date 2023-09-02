output "manifests" {
  value = {
    "example" = data.k8s_sagemaker_services_k8s_aws_notebook_instance_lifecycle_config_v1alpha1_manifest.example.yaml
  }
}
