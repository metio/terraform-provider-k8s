data "k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
