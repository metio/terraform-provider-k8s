output "manifests" {
  value = {
    "example" = data.k8s_imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1_manifest.example.yaml
  }
}
