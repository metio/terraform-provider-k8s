output "manifests" {
  value = {
    "example" = data.k8s_networking_istio_io_proxy_config_v1beta1_manifest.example.yaml
  }
}
