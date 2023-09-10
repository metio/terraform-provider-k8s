data "k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
