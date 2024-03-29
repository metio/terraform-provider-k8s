output "manifests" {
  value = {
    "example" = data.k8s_machineconfiguration_openshift_io_machine_config_pool_v1_manifest.example.yaml
  }
}
