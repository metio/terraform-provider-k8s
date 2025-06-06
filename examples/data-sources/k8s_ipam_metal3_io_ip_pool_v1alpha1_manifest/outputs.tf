output "manifests" {
  value = {
    "example" = data.k8s_ipam_metal3_io_ip_pool_v1alpha1_manifest.example.yaml
  }
}
