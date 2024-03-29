output "manifests" {
  value = {
    "example" = data.k8s_tinkerbell_org_hardware_v1alpha2_manifest.example.yaml
  }
}
