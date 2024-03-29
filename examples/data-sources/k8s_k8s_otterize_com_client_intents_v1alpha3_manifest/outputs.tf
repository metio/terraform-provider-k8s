output "manifests" {
  value = {
    "example" = data.k8s_k8s_otterize_com_client_intents_v1alpha3_manifest.example.yaml
  }
}
