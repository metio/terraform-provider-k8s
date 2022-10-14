output "resources" {
  value = {
    "minimal" = k8s_apps_gitlab_com_git_lab_v1beta1.minimal.yaml
    "example" = k8s_apps_gitlab_com_git_lab_v1beta1.example.yaml
  }
}
