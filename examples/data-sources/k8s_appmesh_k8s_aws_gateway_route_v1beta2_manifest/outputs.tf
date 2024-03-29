output "manifests" {
  value = {
    "example" = data.k8s_appmesh_k8s_aws_gateway_route_v1beta2_manifest.example.yaml
  }
}
