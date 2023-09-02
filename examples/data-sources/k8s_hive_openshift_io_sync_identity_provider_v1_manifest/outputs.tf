output "manifests" {
  value = {
    "example" = data.k8s_hive_openshift_io_sync_identity_provider_v1_manifest.example.yaml
  }
}
