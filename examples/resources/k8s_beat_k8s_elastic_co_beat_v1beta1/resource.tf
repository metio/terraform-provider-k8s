resource "k8s_beat_k8s_elastic_co_beat_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
