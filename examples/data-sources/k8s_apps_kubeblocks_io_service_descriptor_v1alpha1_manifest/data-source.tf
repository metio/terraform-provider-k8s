data "k8s_apps_kubeblocks_io_service_descriptor_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
