resource "k8s_apps_gitlab_com_runner_v1beta2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
