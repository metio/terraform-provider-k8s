output "manifests" {
  value = {
    "example" = data.k8s_keycloak_k8s_reddec_net_keycloak_client_v1alpha1_manifest.example.yaml
  }
}
