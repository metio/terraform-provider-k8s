output "manifests" {
  value = {
    "example" = data.k8s_stunner_l7mp_io_static_service_v1alpha1_manifest.example.yaml
  }
}
