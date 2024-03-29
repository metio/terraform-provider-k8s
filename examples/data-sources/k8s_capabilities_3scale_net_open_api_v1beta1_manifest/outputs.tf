output "manifests" {
  value = {
    "example" = data.k8s_capabilities_3scale_net_open_api_v1beta1_manifest.example.yaml
  }
}
