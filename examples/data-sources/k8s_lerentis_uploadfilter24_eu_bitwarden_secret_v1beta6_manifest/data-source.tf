data "k8s_lerentis_uploadfilter24_eu_bitwarden_secret_v1beta6_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
