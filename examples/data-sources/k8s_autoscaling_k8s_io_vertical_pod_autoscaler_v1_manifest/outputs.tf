output "manifests" {
  value = {
    "example" = data.k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1_manifest.example.yaml
  }
}
