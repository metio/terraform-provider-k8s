output "manifests" {
  value = {
    "example" = data.k8s_tf_tungsten_io_control_v1alpha1_manifest.example.yaml
  }
}
