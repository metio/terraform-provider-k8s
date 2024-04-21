data "k8s_cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
