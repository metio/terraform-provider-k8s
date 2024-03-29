output "manifests" {
  value = {
    "example" = data.k8s_route_openshift_io_route_v1_manifest.example.yaml
  }
}
