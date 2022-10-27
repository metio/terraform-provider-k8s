resource "k8s_apps_clusternet_io_helm_release_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    chart            = "some-helm-chart"
    repo             = "https://example.com/charts"
    target_namespace = "some-namespace"
  }
}
