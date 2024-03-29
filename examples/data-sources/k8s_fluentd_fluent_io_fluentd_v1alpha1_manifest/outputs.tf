output "manifests" {
  value = {
    "example" = data.k8s_fluentd_fluent_io_fluentd_v1alpha1_manifest.example.yaml
  }
}
