output "manifests" {
  value = {
    "example" = data.k8s_apps_deployment_v1_manifest.example.yaml
  }
}
