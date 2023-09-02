output "manifests" {
  value = {
    "example" = data.k8s_logging_banzaicloud_io_logging_v1beta1_manifest.example.yaml
  }
}
