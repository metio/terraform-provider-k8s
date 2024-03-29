data "k8s_snapscheduler_backube_snapshot_schedule_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
