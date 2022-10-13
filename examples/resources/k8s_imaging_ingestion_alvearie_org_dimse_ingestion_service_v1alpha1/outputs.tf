output "resources" {
  value = {
    "minimal" = k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1.minimal.yaml
    "example" = k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1.example.yaml
  }
}
