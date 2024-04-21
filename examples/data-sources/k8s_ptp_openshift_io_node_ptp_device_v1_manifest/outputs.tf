output "manifests" {
  value = {
    "example" = data.k8s_ptp_openshift_io_node_ptp_device_v1_manifest.example.yaml
  }
}
