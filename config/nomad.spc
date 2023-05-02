connection "nomad" {
  plugin = "nomad"

  # `address` - The address of the Nomad server.
  # Can also be set with the NOMAD_ADDR environment variable.
  # address = "http://18.118.164.168:4646"

  # `secret_id` - The SecretID of an ACL token.
  # The SecretID is required to make requests for ACL-enabled clusters.
  # Can also be set with the NOMAD_TOKEN environment variable.
  # secret_id = "c178b810-8b18-6f38-016f-725ddec5d58"
}
