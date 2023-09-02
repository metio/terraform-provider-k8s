output "manifests" {
  value = {
    "example" = data.k8s_metal3_io_bare_metal_host_v1alpha1_manifest.example.yaml
  }
}
