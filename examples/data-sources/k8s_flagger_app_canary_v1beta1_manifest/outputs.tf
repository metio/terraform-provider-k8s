output "manifests" {
  value = {
    "example" = data.k8s_flagger_app_canary_v1beta1_manifest.example.yaml
  }
}
