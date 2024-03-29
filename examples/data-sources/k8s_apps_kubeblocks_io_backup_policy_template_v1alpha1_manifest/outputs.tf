output "manifests" {
  value = {
    "example" = data.k8s_apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest.example.yaml
  }
}
