output "manifests" {
  value = {
    "example" = data.k8s_metacontroller_k8s_io_composite_controller_v1alpha1_manifest.example.yaml
  }
}
