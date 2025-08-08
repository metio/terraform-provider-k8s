data "k8s_csi_ceph_io_ceph_connection_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
