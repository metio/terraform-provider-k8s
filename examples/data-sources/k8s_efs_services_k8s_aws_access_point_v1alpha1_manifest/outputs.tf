output "manifests" {
  value = {
    "example" = data.k8s_efs_services_k8s_aws_access_point_v1alpha1_manifest.example.yaml
  }
}
