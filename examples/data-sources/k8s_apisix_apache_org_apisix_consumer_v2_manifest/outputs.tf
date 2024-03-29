output "manifests" {
  value = {
    "example" = data.k8s_apisix_apache_org_apisix_consumer_v2_manifest.example.yaml
  }
}
