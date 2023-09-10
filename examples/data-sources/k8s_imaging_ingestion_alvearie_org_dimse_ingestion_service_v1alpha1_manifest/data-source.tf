data "k8s_imaging_ingestion_alvearie_org_dimse_ingestion_service_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    application_entity_title : "DICOM-INGEST"
    bucket_config_name : "imaging-ingestion"
    bucket_secret_name : "imaging-ingestion"
    dicom_event_driven_ingestion_name : "core"
    dimse_service : {}
    image_pull_policy : "Always"
    nats_secure : true
    nats_subject_root : "DIMSE"
    nats_token_secret : "ingestion-nats-secure-bound-token"
    nats_url : "nats-secure.imaging-ingestion.svc.cluster.local:4222"
    provider_name : "provider"
  }
}
