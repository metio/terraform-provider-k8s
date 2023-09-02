output "manifests" {
  value = {
    "example" = data.k8s_ec2_services_k8s_aws_route_table_v1alpha1_manifest.example.yaml
  }
}
