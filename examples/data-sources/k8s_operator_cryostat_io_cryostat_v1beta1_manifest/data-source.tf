data "k8s_operator_cryostat_io_cryostat_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    enable_cert_manager  = true
    event_templates      = []
    minimal              = false
    trusted_cert_secrets = []
    report_options = {
      replicas = 0
    }
    storage_options = {
      pvc = {
        annotations = {}
        labels      = {}
        spec        = {}
      }
    }
  }
}
