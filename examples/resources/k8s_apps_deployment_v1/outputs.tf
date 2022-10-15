output "resources" {
  value = {
    "minimal" = k8s_apps_deployment_v1.minimal.yaml
    "example" = k8s_apps_deployment_v1.example.yaml
  }
}
