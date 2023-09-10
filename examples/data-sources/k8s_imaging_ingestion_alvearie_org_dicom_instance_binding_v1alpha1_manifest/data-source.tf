data "k8s_imaging_ingestion_alvearie_org_dicom_instance_binding_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    binding_config_name               = "instance-binding-config"
    binding_secret_name               = "instance-binding-secret"
    dicom_event_driven_ingestion_name = "core"
    image_pull_policy                 = "Always"
    instance_binding = {
      concurrency  = 0
      max_replicas = 3
      min_replicas = 0
    }
  }
}
