resource "k8s_aquasecurity_github_io_aqua_starboard_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
