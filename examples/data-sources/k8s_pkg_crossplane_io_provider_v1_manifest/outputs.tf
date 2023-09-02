output "manifests" {
  value = {
    "example" = data.k8s_pkg_crossplane_io_provider_v1_manifest.example.yaml
  }
}
