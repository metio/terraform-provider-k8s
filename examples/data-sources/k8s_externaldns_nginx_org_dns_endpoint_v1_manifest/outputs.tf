output "manifests" {
  value = {
    "example" = data.k8s_externaldns_nginx_org_dns_endpoint_v1_manifest.example.yaml
  }
}
