data "k8s_jobsmanager_raczylo_com_managed_job_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
