output "manifests" {
  value = {
    "example" = data.k8s_edp_epam_com_sonar_user_v1alpha1_manifest.example.yaml
  }
}
