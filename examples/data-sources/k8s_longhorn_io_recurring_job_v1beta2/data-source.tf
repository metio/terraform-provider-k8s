data "k8s_longhorn_io_recurring_job_v1beta2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
