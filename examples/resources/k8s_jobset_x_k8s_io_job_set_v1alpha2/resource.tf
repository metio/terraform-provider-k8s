resource "k8s_jobset_x_k8s_io_job_set_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
