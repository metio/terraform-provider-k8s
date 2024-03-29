output "manifests" {
  value = {
    "example" = data.k8s_crane_konveyor_io_operator_config_v1alpha1_manifest.example.yaml
  }
}
