output "manifests" {
  value = {
    "example" = data.k8s_fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest.example.yaml
  }
}
