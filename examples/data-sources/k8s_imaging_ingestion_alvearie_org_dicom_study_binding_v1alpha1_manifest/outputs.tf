output "manifests" {
  value = {
    "example" = data.k8s_imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1_manifest.example.yaml
  }
}
