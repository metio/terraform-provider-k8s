resource "k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
