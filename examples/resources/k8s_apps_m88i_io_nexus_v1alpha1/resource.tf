resource "k8s_apps_m88i_io_nexus_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
