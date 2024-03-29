output "manifests" {
  value = {
    "example" = data.k8s_fluentbit_fluent_io_parser_v1alpha2_manifest.example.yaml
  }
}
