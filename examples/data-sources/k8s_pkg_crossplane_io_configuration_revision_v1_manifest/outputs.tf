output "manifests" {
  value = {
    "example" = data.k8s_pkg_crossplane_io_configuration_revision_v1_manifest.example.yaml
  }
}
