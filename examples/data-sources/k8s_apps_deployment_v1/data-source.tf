data "k8s_apps_deployment_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
