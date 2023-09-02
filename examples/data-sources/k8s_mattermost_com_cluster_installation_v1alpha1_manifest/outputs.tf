output "manifests" {
  value = {
    "example" = data.k8s_mattermost_com_cluster_installation_v1alpha1_manifest.example.yaml
  }
}
