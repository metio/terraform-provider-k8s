output "manifests" {
  value = {
    "example" = data.k8s_config_openshift_io_scheduler_v1_manifest.example.yaml
  }
}
