output "manifests" {
  value = {
    "example" = data.k8s_ocmagent_managed_openshift_io_ocm_agent_v1alpha1_manifest.example.yaml
  }
}
