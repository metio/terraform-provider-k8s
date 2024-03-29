output "manifests" {
  value = {
    "example" = data.k8s_apisix_apache_org_apisix_global_rule_v2_manifest.example.yaml
  }
}
