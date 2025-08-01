output "manifests" {
  value = {
    "example" = data.k8s_v2_edp_epam_com_gerrit_project_v1_manifest.example.yaml
  }
}
