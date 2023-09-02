output "manifests" {
  value = {
    "example" = data.k8s_gateway_solo_io_virtual_host_option_v1_manifest.example.yaml
  }
}
