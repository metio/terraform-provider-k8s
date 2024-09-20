data "k8s_listeners_stackable_tech_pod_listeners_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
