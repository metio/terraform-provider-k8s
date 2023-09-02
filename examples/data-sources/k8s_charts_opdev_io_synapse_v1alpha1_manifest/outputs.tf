output "manifests" {
  value = {
    "example" = data.k8s_charts_opdev_io_synapse_v1alpha1_manifest.example.yaml
  }
}
