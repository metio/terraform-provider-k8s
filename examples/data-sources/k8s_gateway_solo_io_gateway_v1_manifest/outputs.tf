output "manifests" {
  value = {
    "example" = data.k8s_gateway_solo_io_gateway_v1_manifest.example.yaml
  }
}
