output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_traffic_route_v1alpha1_manifest.example.yaml
  }
}
