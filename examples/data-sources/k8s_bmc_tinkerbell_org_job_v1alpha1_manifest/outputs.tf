output "manifests" {
  value = {
    "example" = data.k8s_bmc_tinkerbell_org_job_v1alpha1_manifest.example.yaml
  }
}
