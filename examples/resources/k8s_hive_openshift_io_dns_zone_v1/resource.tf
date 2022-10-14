resource "k8s_hive_openshift_io_dns_zone_v1" "minimal" {
  metadata = {
    name = "test"
  }
}
