output "manifests" {
  value = {
    "example" = data.k8s_example_openshift_io_stable_config_type_v1_manifest.example.yaml
  }
}
