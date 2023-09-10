data "k8s_apps_gitlab_com_git_lab_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
