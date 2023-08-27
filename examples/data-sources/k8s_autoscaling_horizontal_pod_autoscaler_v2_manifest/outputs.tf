output "manifests" {
  value = {
    "example" = data.k8s_autoscaling_horizontal_pod_autoscaler_v2_manifest.example.yaml
  }
}
