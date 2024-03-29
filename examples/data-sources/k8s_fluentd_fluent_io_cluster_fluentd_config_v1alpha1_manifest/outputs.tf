output "manifests" {
  value = {
    "example" = data.k8s_fluentd_fluent_io_cluster_fluentd_config_v1alpha1_manifest.example.yaml
  }
}
