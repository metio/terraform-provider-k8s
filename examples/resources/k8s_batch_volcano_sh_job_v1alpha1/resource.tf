resource "k8s_batch_volcano_sh_job_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
