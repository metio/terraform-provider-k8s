data "k8s_hazelcast_com_cron_hot_backup_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}