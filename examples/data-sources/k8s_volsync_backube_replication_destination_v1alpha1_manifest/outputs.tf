output "manifests" {
  value = {
    "example" = data.k8s_volsync_backube_replication_destination_v1alpha1_manifest.example.yaml
  }
}
