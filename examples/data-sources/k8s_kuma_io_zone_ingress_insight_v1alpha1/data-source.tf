data "k8s_kuma_io_zone_ingress_insight_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}