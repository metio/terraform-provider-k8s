data "k8s_imaging_ingestion_alvearie_org_dicom_event_bridge_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    dicom_event_driven_ingestion_name : "core"
    event_bridge : {}
    image_pull_policy : "Always"
    nats_secure : true
    nats_subject_root : "events"
    nats_token_secret : "nats-events-secure-bound-token"
    nats_url : "jetstream.imaging-ingestion.svc.cluster.local:4222"
    role : "hub"
  }
}
