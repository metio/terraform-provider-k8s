data "k8s_apps_kubeedge_io_edge_application_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
