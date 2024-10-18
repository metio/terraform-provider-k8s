output "manifests" {
  value = {
    "example" = data.k8s_kueue_x_k8s_io_multi_kueue_cluster_v1beta1_manifest.example.yaml
  }
}
