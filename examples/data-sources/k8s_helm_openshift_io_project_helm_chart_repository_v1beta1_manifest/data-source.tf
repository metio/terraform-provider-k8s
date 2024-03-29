data "k8s_helm_openshift_io_project_helm_chart_repository_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
