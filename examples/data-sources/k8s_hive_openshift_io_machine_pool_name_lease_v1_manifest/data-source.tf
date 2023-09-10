data "k8s_hive_openshift_io_machine_pool_name_lease_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
