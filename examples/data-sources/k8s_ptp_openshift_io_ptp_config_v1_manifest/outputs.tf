output "manifests" {
  value = {
    "example" = data.k8s_ptp_openshift_io_ptp_config_v1_manifest.example.yaml
  }
}
