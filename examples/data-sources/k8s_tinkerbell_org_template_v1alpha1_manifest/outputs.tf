output "manifests" {
  value = {
    "example" = data.k8s_tinkerbell_org_template_v1alpha1_manifest.example.yaml
  }
}
