output "manifests" {
  value = {
    "example" = data.k8s_appmesh_k8s_aws_virtual_node_v1beta2_manifest.example.yaml
  }
}
