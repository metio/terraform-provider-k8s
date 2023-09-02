resource "k8s_apps_clusternet_io_helm_release_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
