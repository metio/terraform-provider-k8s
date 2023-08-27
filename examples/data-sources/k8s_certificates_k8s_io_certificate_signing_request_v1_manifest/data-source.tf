locals {
  cert_request = <<EOT
-----BEGIN CERTIFICATE REQUEST-----
MIHSMIGBAgEAMCoxGDAWBgNVBAoTD2V4YW1wbGUgY2x1c3RlcjEOMAwGA1UEAxMF
YWRtaW4wTjAQBgcqhkjOPQIBBgUrgQQAIQM6AASSG8S2+hQvfMq5ucngPCzK0m0C
ImigHcF787djpF2QDbz3oQ3QsM/I7ftdjB/HHlG2a5YpqjzT0KAAMAoGCCqGSM49
BAMCA0AAMD0CHQDErNLjX86BVfOsYh/A4zmjmGknZpc2u6/coTHqAhxcR41hEU1I
DpNPvh30e0Js8/DYn2YUfu/pQU19
-----END CERTIFICATE REQUEST-----
EOT
}


data "k8s_certificates_k8s_io_certificate_signing_request_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    usages      = ["client auth", "server auth"]
    signer_name = "kubernetes.io/kube-apiserver-client"
    request     = base64encode(local.cert_request)
  }
}
