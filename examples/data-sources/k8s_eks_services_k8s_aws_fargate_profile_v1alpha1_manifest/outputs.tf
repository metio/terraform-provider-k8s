output "manifests" {
  value = {
    "example" = data.k8s_eks_services_k8s_aws_fargate_profile_v1alpha1_manifest.example.yaml
  }
}
