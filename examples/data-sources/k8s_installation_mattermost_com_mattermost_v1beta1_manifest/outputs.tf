output "manifests" {
  value = {
    "example" = data.k8s_installation_mattermost_com_mattermost_v1beta1_manifest.example.yaml
  }
}
