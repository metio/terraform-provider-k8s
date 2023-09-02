output "manifests" {
  value = {
    "example" = data.k8s_traefik_io_middleware_v1alpha1_manifest.example.yaml
  }
}
