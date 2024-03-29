output "manifests" {
  value = {
    "example" = data.k8s_ray_io_ray_service_v1_manifest.example.yaml
  }
}
