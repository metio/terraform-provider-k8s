output "manifests" {
  value = {
    "example" = data.k8s_apps_gitlab_com_git_lab_v1beta1_manifest.example.yaml
  }
}
