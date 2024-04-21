output "manifests" {
  value = {
    "example" = data.k8s_control_k8ssandra_io_cassandra_task_v1alpha1_manifest.example.yaml
  }
}
