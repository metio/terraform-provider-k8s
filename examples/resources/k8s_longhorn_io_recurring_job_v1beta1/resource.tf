resource "k8s_longhorn_io_recurring_job_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
