output "resources" {
  value = {
    "minimal" = k8s_ceph_rook_io_ceph_bucket_notification_v1.minimal.yaml
  }
}
