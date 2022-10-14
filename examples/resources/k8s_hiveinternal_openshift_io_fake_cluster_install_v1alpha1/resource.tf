resource "k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    cluster_deployment_ref = {
      name = "some-deployment"
    }
    image_set_ref = {
      name = "some-image-set"
    }
  }
}
