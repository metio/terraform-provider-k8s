output "manifests" {
  value = {
    "example" = data.k8s_capsule_clastix_io_tenant_v1beta2_manifest.example.yaml
  }
}
