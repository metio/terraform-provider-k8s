output "manifests" {
  value = {
    "example" = data.k8s_application_networking_k8s_aws_target_group_policy_v1alpha1_manifest.example.yaml
  }
}
