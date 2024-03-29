output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_egress_gateway_v1_manifest.example.yaml
  }
}
