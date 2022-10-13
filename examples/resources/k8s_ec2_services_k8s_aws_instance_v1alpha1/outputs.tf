output "resources" {
  value = {
    "minimal" = k8s_ec2_services_k8s_aws_instance_v1alpha1.minimal.yaml
  }
}
