output "manifests" {
  value = {
    "example" = data.k8s_cluster_x_k8s_io_cluster_v1alpha3_manifest.example.yaml
  }
}
