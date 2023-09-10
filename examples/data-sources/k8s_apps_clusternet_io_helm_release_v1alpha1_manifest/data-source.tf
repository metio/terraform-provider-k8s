data "k8s_apps_clusternet_io_helm_release_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    chart            = "some-helm-chart"
    repo             = "https://example.com/charts"
    target_namespace = "some-namespace"
  }
}
