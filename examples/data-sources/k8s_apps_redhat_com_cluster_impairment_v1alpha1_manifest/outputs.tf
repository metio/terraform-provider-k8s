output "manifests" {
  value = {
    "example" = data.k8s_apps_redhat_com_cluster_impairment_v1alpha1_manifest.example.yaml
  }
}
