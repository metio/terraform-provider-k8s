output "manifests" {
  value = {
    "example" = data.k8s_apps_kubedl_io_cron_v1alpha1_manifest.example.yaml
  }
}
