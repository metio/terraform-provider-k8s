output "manifests" {
  value = {
    "example" = data.k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest.example.yaml
  }
}
