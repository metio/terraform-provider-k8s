data "k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
