output "manifests" {
  value = {
    "example" = data.k8s_reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1_manifest.example.yaml
  }
}
