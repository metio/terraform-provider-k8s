output "manifests" {
  value = {
    "example" = data.k8s_bmc_tinkerbell_org_machine_v1alpha1_manifest.example.yaml
  }
}
