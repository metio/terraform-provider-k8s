output "manifests" {
  value = {
    "example" = data.k8s_projectcontour_io_http_proxy_v1_manifest.example.yaml
  }
}
