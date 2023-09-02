output "manifests" {
  value = {
    "example" = data.k8s_sparkoperator_k8s_io_spark_application_v1beta2_manifest.example.yaml
  }
}
