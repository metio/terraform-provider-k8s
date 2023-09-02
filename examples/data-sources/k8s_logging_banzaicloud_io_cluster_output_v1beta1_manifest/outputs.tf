output "manifests" {
  value = {
    "example" = data.k8s_logging_banzaicloud_io_cluster_output_v1beta1_manifest.example.yaml
  }
}
