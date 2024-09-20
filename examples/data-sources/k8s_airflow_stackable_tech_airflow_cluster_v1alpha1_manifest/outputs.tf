output "manifests" {
  value = {
    "example" = data.k8s_airflow_stackable_tech_airflow_cluster_v1alpha1_manifest.example.yaml
  }
}
