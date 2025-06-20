output "manifests" {
  value = {
    "example" = data.k8s_gateway_nginx_org_observability_policy_v1alpha2_manifest.example.yaml
  }
}
