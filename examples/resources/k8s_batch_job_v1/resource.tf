resource "k8s_batch_job_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
