data "k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
