resource "k8s_argoproj_io_application_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    project = "some-project"
    source = {
      repo_url = "https://example.com/repo.git"
    }
    destination = {

    }
  }
}
