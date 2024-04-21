output "manifests" {
  value = {
    "example" = data.k8s_storageos_com_storage_os_cluster_v1_manifest.example.yaml
  }
}
