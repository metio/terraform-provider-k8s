output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_felix_configuration_v1_manifest.example.yaml
  }
}
