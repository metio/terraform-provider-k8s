output "manifests" {
  value = {
    "example" = data.k8s_apps_3scale_net_api_manager_restore_v1alpha1_manifest.example.yaml
  }
}
