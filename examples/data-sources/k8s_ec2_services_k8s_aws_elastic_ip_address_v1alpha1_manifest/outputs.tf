output "manifests" {
  value = {
    "example" = data.k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1_manifest.example.yaml
  }
}
