data "k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    application_entity_title = "DCM4CHEE"
    image_pull_policy        = "Always"
    nats_secure              = true
    nats_subject_root        = "DIMSE"
    nats_token_secret        = "ingestion-nats-secure-bound-token"
    nats_url                 = "nats-secure.imaging-ingestion.svc.cluster.local:4222"
    proxy                    = {}
    target_dimse_host        = "arc.dcm4chee.svc.cluster.local"
    target_dimse_port        = 11112
  }
}
