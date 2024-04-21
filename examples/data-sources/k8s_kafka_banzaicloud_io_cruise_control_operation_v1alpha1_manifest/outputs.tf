output "manifests" {
  value = {
    "example" = data.k8s_kafka_banzaicloud_io_cruise_control_operation_v1alpha1_manifest.example.yaml
  }
}
