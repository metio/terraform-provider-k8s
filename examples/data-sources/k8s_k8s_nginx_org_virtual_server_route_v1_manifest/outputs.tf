output "manifests" {
  value = {
    "example" = data.k8s_k8s_nginx_org_virtual_server_route_v1_manifest.example.yaml
  }
}
