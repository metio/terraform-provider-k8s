output "manifests" {
  value = {
    "example" = data.k8s_v2_edp_epam_com_template_v1alpha1_manifest.example.yaml
  }
}
