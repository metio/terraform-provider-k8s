resource "k8s_operator_cryostat_io_cryostat_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_operator_cryostat_io_cryostat_v1beta1" "example" {
  metadata = {
    name = "cryostat-sample"
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
