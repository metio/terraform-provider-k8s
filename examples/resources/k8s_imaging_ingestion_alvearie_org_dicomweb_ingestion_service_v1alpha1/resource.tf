resource "k8s_imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_imaging_ingestion_alvearie_org_dicomweb_ingestion_service_v1alpha1" "example" {
  metadata = {
    name = "ingestion"
  }
  spec = {
    bucket_config_name                = "imaging-ingestion"
    bucket_secret_name                = "imaging-ingestion"
    dicom_event_driven_ingestion_name = "core"
    image_pull_policy                 = "Always"
    provider_name                     = "provider"
    stow_service = {
      concurrency  = 0
      max_replicas = 3
      min_replicas = 0
    }
    wado_service = {
      concurrency  = 0
      max_replicas = 3
      min_replicas = 0
    }
  }
}
