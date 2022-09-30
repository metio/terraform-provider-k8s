resource "k8s_app_lightbend_com_akka_cluster_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
