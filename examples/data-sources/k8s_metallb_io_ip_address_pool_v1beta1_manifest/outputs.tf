output "manifests" {
  value = {
    "example" = data.k8s_metallb_io_ip_address_pool_v1beta1_manifest.example.yaml
  }
}
