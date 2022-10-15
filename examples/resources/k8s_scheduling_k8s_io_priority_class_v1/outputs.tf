output "resources" {
  value = {
    "minimal" = k8s_scheduling_k8s_io_priority_class_v1.minimal.yaml
    "example" = k8s_scheduling_k8s_io_priority_class_v1.example.yaml
  }
}
