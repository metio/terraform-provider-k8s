data "k8s_argoproj_io_application_set_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    generators = []
    template = {
      metadata = {
        name = "template"
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
  }
}
