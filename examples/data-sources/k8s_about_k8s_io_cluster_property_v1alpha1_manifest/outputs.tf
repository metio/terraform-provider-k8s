output "manifests" {
  value = {
    "example" = data.k8s_about_k8s_io_cluster_property_v1alpha1_manifest.example.yaml
  }
}
