output "resources" {
  value = {
    "minimal" = k8s_autoscaling_horizontal_pod_autoscaler_v2.minimal.yaml
    "example" = k8s_autoscaling_horizontal_pod_autoscaler_v2.example.yaml
  }
}
