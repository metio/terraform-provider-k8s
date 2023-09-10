data "k8s_apps_gitlab_com_runner_v1beta2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
