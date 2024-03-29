output "manifests" {
  value = {
    "example" = data.k8s_tinkerbell_org_osie_v1alpha2_manifest.example.yaml
  }
}
