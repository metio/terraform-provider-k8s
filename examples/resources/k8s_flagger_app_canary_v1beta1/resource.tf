resource "k8s_flagger_app_canary_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
