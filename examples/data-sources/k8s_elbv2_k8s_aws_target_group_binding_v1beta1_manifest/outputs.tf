output "manifests" {
  value = {
    "example" = data.k8s_elbv2_k8s_aws_target_group_binding_v1beta1_manifest.example.yaml
  }
}
