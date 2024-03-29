output "manifests" {
  value = {
    "example" = data.k8s_chaosblade_io_chaos_blade_v1alpha1_manifest.example.yaml
  }
}
