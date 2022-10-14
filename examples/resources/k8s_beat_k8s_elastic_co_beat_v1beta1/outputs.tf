output "resources" {
  value = {
    "minimal" = k8s_beat_k8s_elastic_co_beat_v1beta1.minimal.yaml
    "example" = k8s_beat_k8s_elastic_co_beat_v1beta1.example.yaml
  }
}
