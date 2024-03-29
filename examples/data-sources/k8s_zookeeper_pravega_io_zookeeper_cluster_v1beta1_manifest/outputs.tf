output "manifests" {
  value = {
    "example" = data.k8s_zookeeper_pravega_io_zookeeper_cluster_v1beta1_manifest.example.yaml
  }
}
