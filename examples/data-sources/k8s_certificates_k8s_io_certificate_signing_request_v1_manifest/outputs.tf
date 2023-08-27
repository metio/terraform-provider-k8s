output "manifests" {
  value = {
    "example" = data.k8s_certificates_k8s_io_certificate_signing_request_v1_manifest.example.yaml
  }
}
