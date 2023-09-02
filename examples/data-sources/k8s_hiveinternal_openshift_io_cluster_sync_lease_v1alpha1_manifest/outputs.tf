output "manifests" {
  value = {
    "example" = data.k8s_hiveinternal_openshift_io_cluster_sync_lease_v1alpha1_manifest.example.yaml
  }
}
