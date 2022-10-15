output "resources" {
  value = {
    "minimal" = k8s_certificates_k8s_io_certificate_signing_request_v1.minimal.yaml
    "example" = k8s_certificates_k8s_io_certificate_signing_request_v1.example.yaml
  }
}
