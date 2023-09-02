data "k8s_hive_openshift_io_cluster_relocate_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
