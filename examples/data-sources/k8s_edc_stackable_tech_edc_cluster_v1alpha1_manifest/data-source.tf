data "k8s_edc_stackable_tech_edc_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      cert_secret = "some-secret"
      ionos = {
        s3           = {}
        token_secret = "some-token"
      }
    }
    image = {}
  }
}
