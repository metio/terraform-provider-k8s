output "manifests" {
  value = {
    "example" = data.k8s_logging_banzaicloud_io_syslog_ng_flow_v1beta1_manifest.example.yaml
  }
}
