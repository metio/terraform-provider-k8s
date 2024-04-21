output "manifests" {
  value = {
    "example" = data.k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest.example.yaml
  }
}
