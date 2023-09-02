output "manifests" {
  value = {
    "example" = data.k8s_keycloak_org_keycloak_client_v1alpha1_manifest.example.yaml
  }
}
