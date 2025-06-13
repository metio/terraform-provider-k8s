output "manifests" {
  value = {
    "example" = data.k8s_gateway_networking_x_k8s_io_x_backend_traffic_policy_v1alpha1_manifest.example.yaml
  }
}
