output "manifests" {
  value = {
    "example" = data.k8s_upgrade_managed_openshift_io_upgrade_config_v1alpha1_manifest.example.yaml
  }
}
