output "manifests" {
  value = {
    "example" = data.k8s_prometheusservice_services_k8s_aws_logging_configuration_v1alpha1_manifest.example.yaml
  }
}
