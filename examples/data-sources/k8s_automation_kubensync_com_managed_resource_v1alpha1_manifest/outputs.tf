output "manifests" {
  value = {
    "example" = data.k8s_automation_kubensync_com_managed_resource_v1alpha1_manifest.example.yaml
  }
}
