output "manifests" {
  value = {
    "example" = data.k8s_eks_services_k8s_aws_nodegroup_v1alpha1_manifest.example.yaml
  }
}
