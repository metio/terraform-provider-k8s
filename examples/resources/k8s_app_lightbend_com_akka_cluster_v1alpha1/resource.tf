resource "k8s_app_lightbend_com_akka_cluster_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
