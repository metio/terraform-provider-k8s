data "k8s_hive_openshift_io_dns_zone_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
