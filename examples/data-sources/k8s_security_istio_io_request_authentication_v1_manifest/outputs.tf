output "manifests" {
  value = {
    "example" = data.k8s_security_istio_io_request_authentication_v1_manifest.example.yaml
  }
}
