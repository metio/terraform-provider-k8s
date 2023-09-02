output "manifests" {
  value = {
    "example" = data.k8s_beat_k8s_elastic_co_beat_v1beta1_manifest.example.yaml
  }
}
