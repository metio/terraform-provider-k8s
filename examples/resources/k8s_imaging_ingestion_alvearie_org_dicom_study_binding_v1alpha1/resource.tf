resource "k8s_imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1" "example" {
  metadata = {
    name = "fhir"
  }
  spec = {
    binding_config_name               = "study-binding-config"
    binding_secret_name               = "study-binding-secret"
    dicom_event_driven_ingestion_name = "core"
    image_pull_policy                 = "Always"
    study_binding = {
      concurrency  = 0
      max_replicas = 3
      min_replicas = 0
    }
  }
}
