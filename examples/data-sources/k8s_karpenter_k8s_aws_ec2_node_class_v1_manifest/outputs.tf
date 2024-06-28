output "manifests" {
  value = {
    "example" = data.k8s_karpenter_k8s_aws_ec2_node_class_v1_manifest.example.yaml
  }
}
