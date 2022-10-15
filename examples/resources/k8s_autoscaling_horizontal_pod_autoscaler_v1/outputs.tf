output "resources" {
  value = {
    "minimal" = k8s_autoscaling_horizontal_pod_autoscaler_v1.minimal.yaml
    "example" = k8s_autoscaling_horizontal_pod_autoscaler_v1.example.yaml
  }
}
