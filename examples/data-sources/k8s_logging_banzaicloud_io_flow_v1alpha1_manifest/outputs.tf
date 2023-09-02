output "manifests" {
  value = {
    "example" = data.k8s_logging_banzaicloud_io_flow_v1alpha1_manifest.example.yaml
  }
}
