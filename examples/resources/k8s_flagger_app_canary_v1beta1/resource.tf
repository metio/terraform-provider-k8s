resource "k8s_flagger_app_canary_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
