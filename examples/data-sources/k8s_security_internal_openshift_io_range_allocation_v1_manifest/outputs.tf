output "manifests" {
  value = {
    "example" = data.k8s_security_internal_openshift_io_range_allocation_v1_manifest.example.yaml
  }
}
