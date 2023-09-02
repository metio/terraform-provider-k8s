output "manifests" {
  value = {
    "example" = data.k8s_data_fluid_io_thin_runtime_profile_v1alpha1_manifest.example.yaml
  }
}
