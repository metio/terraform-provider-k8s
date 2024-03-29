output "manifests" {
  value = {
    "example" = data.k8s_vpcresources_k8s_aws_cni_node_v1alpha1_manifest.example.yaml
  }
}
