data "k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
