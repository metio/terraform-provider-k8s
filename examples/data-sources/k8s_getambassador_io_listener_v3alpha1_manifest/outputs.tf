output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_listener_v3alpha1_manifest.example.yaml
  }
}
