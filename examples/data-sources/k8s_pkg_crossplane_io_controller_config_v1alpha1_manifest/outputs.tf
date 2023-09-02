output "manifests" {
  value = {
    "example" = data.k8s_pkg_crossplane_io_controller_config_v1alpha1_manifest.example.yaml
  }
}
