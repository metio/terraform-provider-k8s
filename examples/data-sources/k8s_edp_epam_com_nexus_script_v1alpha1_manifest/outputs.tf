output "manifests" {
  value = {
    "example" = data.k8s_edp_epam_com_nexus_script_v1alpha1_manifest.example.yaml
  }
}
