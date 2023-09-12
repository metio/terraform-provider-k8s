output "manifests" {
  value = {
    "example" = data.k8s_hnc_x_k8s_io_hnc_configuration_v1alpha2_manifest.example.yaml
  }
}
