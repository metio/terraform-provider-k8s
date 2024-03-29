output "manifests" {
  value = {
    "example" = data.k8s_stunner_l7mp_io_dataplane_v1alpha1_manifest.example.yaml
  }
}
