output "manifests" {
  value = {
    "example" = data.k8s_clusters_clusternet_io_cluster_registration_request_v1beta1_manifest.example.yaml
  }
}
