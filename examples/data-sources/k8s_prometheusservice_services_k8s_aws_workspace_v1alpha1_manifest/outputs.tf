output "manifests" {
  value = {
    "example" = data.k8s_prometheusservice_services_k8s_aws_workspace_v1alpha1_manifest.example.yaml
  }
}
