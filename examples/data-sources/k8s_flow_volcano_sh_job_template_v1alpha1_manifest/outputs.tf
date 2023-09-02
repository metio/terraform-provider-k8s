output "manifests" {
  value = {
    "example" = data.k8s_flow_volcano_sh_job_template_v1alpha1_manifest.example.yaml
  }
}
