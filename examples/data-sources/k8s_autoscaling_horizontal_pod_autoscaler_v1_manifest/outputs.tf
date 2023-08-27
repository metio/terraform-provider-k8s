output "manifests" {
  value = {
    "example" = data.k8s_autoscaling_horizontal_pod_autoscaler_v1_manifest.example.yaml
  }
}
