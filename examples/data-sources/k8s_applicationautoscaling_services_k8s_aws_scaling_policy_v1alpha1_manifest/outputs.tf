output "manifests" {
  value = {
    "example" = data.k8s_applicationautoscaling_services_k8s_aws_scaling_policy_v1alpha1_manifest.example.yaml
  }
}
