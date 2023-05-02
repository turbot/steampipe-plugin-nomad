connection "nomad" {
  plugin = "nomad"

  # `address` - The address of the Nomad server.
  # Can also be set with the NOMAD_ADDR environment variable.
  # address = "http://18.118.164.168:4646"

  # `namespace` - The Nomad cluster namespace.
  # For more information on the Namespace, please see https://developer.hashicorp.com/nomad/tutorials/manage-clusters/namespaces.
  # Can also be set with the NOMAD_NAMESPACE environment variable.
  # By default, the plugin will query all the namespaces.
  namespace = "*"

  # `secret_id` - The SecretID of an ACL token.
  # The SecretID is required to make requests for ACL-enabled clusters.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/nomad/tutorials/access-control/access-control-tokens.
  # Can also be set with the NOMAD_TOKEN environment variable.
  # secret_id = "c178b810-8b18-6f38-016f-725ddec5d58"
}
