output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_tls_pass_through_route_v1_manifest.example.yaml
  }
}
