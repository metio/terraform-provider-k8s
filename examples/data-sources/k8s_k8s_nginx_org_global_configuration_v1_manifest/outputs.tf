output "manifests" {
  value = {
    "example" = data.k8s_k8s_nginx_org_global_configuration_v1_manifest.example.yaml
  }
}
