resource "k8s_argoproj_io_app_project_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {

  }
}
