resource "k8s_helm_sigstore_dev_rekor_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
