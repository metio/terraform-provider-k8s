output "manifests" {
  value = {
    "example" = data.k8s_cluster_x_k8s_io_cluster_class_v1beta1_manifest.example.yaml
  }
}
