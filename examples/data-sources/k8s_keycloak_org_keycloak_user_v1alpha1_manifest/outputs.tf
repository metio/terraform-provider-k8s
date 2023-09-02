output "manifests" {
  value = {
    "example" = data.k8s_keycloak_org_keycloak_user_v1alpha1_manifest.example.yaml
  }
}
