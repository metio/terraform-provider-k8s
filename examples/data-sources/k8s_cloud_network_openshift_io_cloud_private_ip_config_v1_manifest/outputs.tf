output "manifests" {
  value = {
    "example" = data.k8s_cloud_network_openshift_io_cloud_private_ip_config_v1_manifest.example.yaml
  }
}
