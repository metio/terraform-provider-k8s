data "k8s_imaging_ingestion_alvearie_org_dicom_event_driven_ingestion_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    database_config_name = "db-config"
    database_secret_name = "db-secret"
    event_processor = {
      concurrency  = 0
      max_replicas = 3
      min_replicas = 0
    }
    image_pull_policy = "Always"
  }
}
