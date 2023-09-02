output "manifests" {
  value = {
    "example" = data.k8s_gateway_nginx_org_nginx_gateway_v1alpha1_manifest.example.yaml
  }
}
