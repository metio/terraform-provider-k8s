output "manifests" {
  value = {
    "example" = data.k8s_data_fluid_io_alluxio_runtime_v1alpha1_manifest.example.yaml
  }
}
