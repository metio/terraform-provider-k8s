output "manifests" {
  value = {
    "example" = data.k8s_vpcresources_k8s_aws_security_group_policy_v1beta1_manifest.example.yaml
  }
}
