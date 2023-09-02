output "manifests" {
  value = {
    "example" = data.k8s_capsule_clastix_io_tenant_v1beta1_manifest.example.yaml
  }
}
