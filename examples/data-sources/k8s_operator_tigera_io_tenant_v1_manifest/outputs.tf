output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_tenant_v1_manifest.example.yaml
  }
}
