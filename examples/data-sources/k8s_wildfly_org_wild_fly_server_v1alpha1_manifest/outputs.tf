output "manifests" {
  value = {
    "example" = data.k8s_wildfly_org_wild_fly_server_v1alpha1_manifest.example.yaml
  }
}
