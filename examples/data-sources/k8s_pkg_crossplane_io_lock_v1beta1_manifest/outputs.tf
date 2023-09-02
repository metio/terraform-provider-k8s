output "manifests" {
  value = {
    "example" = data.k8s_pkg_crossplane_io_lock_v1beta1_manifest.example.yaml
  }
}
