output "manifests" {
  value = {
    "example" = data.k8s_gateway_networking_k8s_io_backend_tls_policy_v1alpha3_manifest.example.yaml
  }
}
