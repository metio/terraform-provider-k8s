output "manifests" {
  value = {
    "example" = data.k8s_cluster_clusterpedia_io_pedia_cluster_v1alpha2_manifest.example.yaml
  }
}
