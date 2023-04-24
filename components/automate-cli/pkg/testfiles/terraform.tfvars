backup_config_s3 = "true"
# Existing nodes
################################################################################
existing_automate_private_ips = ["10.1.0.101"]
existing_chef_server_private_ips = ["10.1.0.102"]
existing_opensearch_private_ips = ["10.1.0.103", "10.1.1.104", "10.1.2.104"]
existing_postgresql_private_ips = ["10.1.0.105", "10.1.1.106", "10.1.2.107"]
setup_managed_services = "false"
setup_self_managed_services = "false"
managed_opensearch_domain_name = ""
managed_opensearch_domain_url = ""
managed_opensearch_username = ""
managed_opensearch_user_password = ""
managed_opensearch_certificate = " "
aws_os_snapshot_role_arn = " "
os_snapshot_user_access_key_id = " "
os_snapshot_user_access_key_secret = " "
managed_rds_instance_url = ""
managed_rds_superuser_username = ""
managed_rds_superuser_password = ""
managed_rds_dbuser_username = ""
managed_rds_dbuser_password = ""
postgresql_root_cert = <<-EOT
EOT
opensearch_root_cert = <<-EOT
EOT
# Common
################################################################################
automate_admin_password = "chefautomate"
automate_config_file = "/hab/a2_deploy_workspace/configs/automate.toml"
automate_fqdn = "A2-9548100b-automate-lb-1013582932.ap-south-1.elb.amazonaws.com" # leave commented out for AWS, othewise must be assigned
automate_instance_count = 1
chef_server_instance_count = 1
opensearch_instance_count = 3
automate_root_ca = <<-EOT
-----START-----
MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8fF5Q=
-----END-----
EOT
opensearch_root_ca = <<-EOT
-----START-----
MIIDQjCCAioCCQCY6unlpgYfsTANBgkqhkiG9w0BAQsFADBjMQswCQYDVQQGEwJV
UzETMBEGA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UE
CgwRQ2hlZiBTb2Z0d2FyZSBJbmMxETAPBgNVBAMMCHByb2dyZXNzMB4XDTIyMTEw
DWvLPsLlfQB2MQl4/oeTHxpjmc/7njK+surbRLznGPQmy1gHML0nVzDSey949YNb
CNoh2WWpCTrNQypFSdzNA4RVcnDwpgULH90+zyHGNLBkarddi1WP1fabq7MnM9lw
7VbwQVw4sIOgs0jYff/gNolVypnjgkWJJBoMcqhCOGGKYj4afuP1x0ibxd2XfwOr
K15v3u9n2qgsXcR7Og2xWeTKpKxg8A==
-----END-----
EOT
postgresql_root_ca = <<-EOT
-----START-----
MIIDQjCCAioCCQCY6unlpgYfsTANBgkqhkiG9w0BAQsFADBjMQswCQYDVQQGEwJV
UzETMBEGA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UE
CgwRQ2hlZiBTb2Z0d2FyZSBJbmMxETAPBgNVBAMMCHByb2dyZXNzMB4XDTIyMTEw
DWvLPsLlfQB2MQl4/oeTHxpjmc/7njK+surbRLznGPQmy1gHML0nVzDSey949YNb
CNoh2WWpCTrNQypFSdzNA4RVcnDwpgULH90+zyHGNLBkarddi1WP1fabq7MnM9lw
7VbwQVw4sIOgs0jYff/gNolVypnjgkWJJBoMcqhCOGGKYj4afuP1x0ibxd2XfwOr
K15v3u9n2qgsXcR7Og2xWeTKpKxg8A==
-----END-----
EOT
automate_private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
BOKnt5GgdKukzw2lBadXL/eF0yg8TNdmgRg1BsYZfFLWZyZ5FE2oqqKyOTWi2M5i
PbTvxWvIP4K8K8b1wnxLcFuN5OtJ21q+1MccprdfVOTMXg3d9ne2LDBkIPYshL+2
P3PKJm6ElG3IIcsi6fLFsF6bq96Qxzzn5udfE4Z1bueNL+NBBq7fDTObMp22SM2V
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
chef_server_private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
BOKnt5GgdKukzw2lBadXL/eF0yg8TNdmgRg1BsYZfFLWZyZ5FE2oqqKyOTWi2M5i
PbTvxWvIP4K8K8b1wnxLcFuN5OtJ21q+1MccprdfVOTMXg3d9ne2LDBkIPYshL+2
hgQ0oX3QuVS3cGdvC0yqVYuQwbxsObEYtwG8mO2pgpOUwGqPlrlKoYx6nBqDo4eq
YvXmWYr7l3Q+TOOdmSxVPv0EoE4BnEC/eL0cj4dx4JwG84fNulbYhgqTdQKBgQC1
P3PKJm6ElG3IIcsi6fLFsF6bq96Qxzzn5udfE4Z1bueNL+NBBq7fDTObMp22SM2V
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
opensearch_admin_key = <<-EOT
-----START-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDfRrez3MbQzX2x
ShyAeJPIKNZRJzoYDLxmuTFIFnJlGTM1/DxIWokMxN/zY5xCjp0cqBMp4MkNjLk7
qbBF8AQADyhHCKnRA2+qnj3amvnJBIwepSCy/CT1jKTpgKNe8ULpsOA/chjH9xUY
6iWLNZqHF9TB6GNJuAmJk/vQeOBCcVhIwla7LHLDl1AZqtbrVvOebgBlyIZE6KRc
liHRb5t7CinFyHpJA6N7bEkdt4O1nariz0x9VFJAAKlj8m4IGzYuIPDRTJzlafJ7
ZcHRwhISaehzOx06iWo2ME1Ca3IpBiMlVAy+xRulZwKBgBq43BjPxU55ubPdEHBK
dwFhHD8n8m6XvFNM3ypzd3j/E8pECfBrN1CJM8WG5ElW7JvST9OPpG7NxaFrXn0f
M2U1kP/54OuAjJapsgz76TLI59EtWUY+2VQxYYh914jywTNBcRlSTE7kAplddMSt
YP35CayWktERhGKoqbGTTNtn
-----END-----
EOT
opensearch_private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
BOKnt5GgdKukzw2lBadXL/eF0yg8TNdmgRg1BsYZfFLWZyZ5FE2oqqKyOTWi2M5i
PbTvxWvIP4K8K8b1wnxLcFuN5OtJ21q+1MccprdfVOTMXg3d9ne2LDBkIPYshL+2
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
postgresql_private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
bWMZsMyWMwKBgQDeMy65IzxdTaew4IScHoPMXY98F5D6l+uHW5erfYH//kH4o8qO
hgQ0oX3QuVS3cGdvC0yqVYuQwbxsObEYtwG8mO2pgpOUwGqPlrlKoYx6nBqDo4eq
YvXmWYr7l3Q+TOOdmSxVPv0EoE4BnEC/eL0cj4dx4JwG84fNulbYhgqTdQKBgQC1
P3PKJm6ElG3IIcsi6fLFsF6bq96Qxzzn5udfE4Z1bueNL+NBBq7fDTObMp22SM2V
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
automate_public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
chef_server_public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
dFNpxiedGkBWqtwJHYNu+kfbHETujMZAb80JYC3Q0MaBKKz1oAQ8yDu6YXHCjgIa
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
opensearch_public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
opensearch_admin_cert = <<-EOT
-----START-----
MIIDazCCAlOgAwIBAgIJAPHEYb1O7wznMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
QUoHCIPqXdbKxPiCVETs+VuLWzrmpShS93VgOzWBEuPMRQJWEVbCMUgc/XLYfdft
6rvtVqBuLgB/7o3KBCF9LoNPPON3sd49Xhn5aL8Hh3vOHuFphPXzhgkdofy499Fp
8eYa8nXTsAV2NwdySnV0zR9BTNO/7G4lFYYKX45REQS1taRFN3yNOMjeD8TuMSFc
z8zNIe7n62slIsQ5jnkaEI32llgrGXWTrx3bogQQubvz9iJacecRJwXuuIigJSVM
1Vz0Nr1NGEP0guArf7Nz
-----END-----
EOT
postgresql_public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
hXIQE/2dUfwC3YZvQlB8NOGoLIAq2MRZZW/efPtsUnf6VMtUUSGGYKaKd0xLJ3YO
l5JX+bjhfauAZqYmaRhkvwPzrWPMhw/wnlkJ4WcNN6OWRwIDAQABoyEwHzAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBALB9
dFNpxiedGkBWqtwJHYNu+kfbHETujMZAb80JYC3Q0MaBKKz1oAQ8yDu6YXHCjgIa
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
automate_custom_certs_enabled = true
chef_server_custom_certs_enabled = true
postgresql_custom_certs_enabled = true
opensearch_custom_certs_enabled = true
opensearch_admin_dn = "CN=chefadmin,O=Chef Software Inc,L=Seattle,ST=Washington,C=US"
opensearch_nodes_dn = "CN=chefnode,O=Chef Software Inc,L=Seattle,ST=Washington,C=US"
automate_certs_by_ip = {  "10.1.0.108" = { private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
RMyJRaCbwi6sOhV2JUdi3m3XIM+MNXrIl5XPWzOk9p2hMjIE+VqpwKNBT7gfucxu
2RjhksDH7p7mbFap0jkUZJ+KKS/KVTEmkXpAW+zZVEWsmFoye2YhxPOw7uQRpkpw
bWMZsMyWMwKBgQDeMy65IzxdTaew4IScHoPMXY98F5D6l+uHW5erfYH//kH4o8qO
hgQ0oX3QuVS3cGdvC0yqVYuQwbxsObEYtwG8mO2pgpOUwGqPlrlKoYx6nBqDo4eq
YvXmWYr7l3Q+TOOdmSxVPv0EoE4BnEC/eL0cj4dx4JwG84fNulbYhgqTdQKBgQC1
P3PKJm6ElG3IIcsi6fLFsF6bq96Qxzzn5udfE4Z1bueNL+NBBq7fDTObMp22SM2V
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBALB9
dFNpxiedGkBWqtwJHYNu+kfbHETujMZAb80JYC3Q0MaBKKz1oAQ8yDu6YXHCjgIa
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
 }  }
chef_server_certs_by_ip = {  "10.1.0.14" = { private_key = <<-EOT
-----START-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC1yT32Ii+EUIAe
FZQbOxzGduI4I6NDicF0VTa22hW++Fdn7lrgq/CirwjlQNKXsImG+9zlQHVIuna3
blXEu7TsLqAb08eAA1PEWhLmasCNUq8SPP8xfndviTFwvNUTneE4WkJ/uQKBgF6V
AUuNLWgBclV6TzYlm10Wo/FHcBSHyl1SOcgkTAdeCKRIm1Jd+T+LIUFcvtPOPMNG
4UAvlM496ruqJnhAAkXuHmdN1KW7zbNw/FW7PnHPQM7RTzkdgHmh3PQHRojZnTN6
ABNejk5jj31GwRc7Qm5JYALhq/GNs8YPsmIATiD7AoGAJQ1N36pQgYy5Jy5I0aIq
p3nGfRWBLYqftL2oyDQGeyRQX9/hF8hh6AwbET/CAPoV+e4EjOI9VeJVjWYZ1raB
CHBWACozIc+2fQQngj5PsH6PwP+sq47BcA8mEorJcYRelKZF1NJp9KSrcaDSo8MJ
w/qTq+SCthUyjrVT50EN34k=
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzqMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
QIzkWVvgj5JgtfHpmZ8XPBYLfl5ZcL5AxtM24TCvhwmghQIDAQABoyEwHzAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAH1a
5ER8bgk1E3600ak0y01E9Li3GbECwshLlx9OH4NesBuFYFuzI5gYd+pKQVnFsRiP
X8CJ0GLWrbfuOk27mIKkKFsuT44WQQMXWb4uf0h7oRENG8ASeiZHw9BmMUCPpufh
6WA1WxtHm4Bq+l3910YMaNZFZi953BUXF6OWfrMYelqlGEDR4UIkFGw5cmHBQJbN
Uw4+ATYf15lE3CQsKDknuRSJYun7M4+Wc8JKArtTPgVJ/la3Y1tpaKf36UL3qkxK
d0XhuTq+G0/vyjRCjAESDkF2Djw2Eb7VFnDPQCfRkKjjgIWjps0/kBVP6T4flQGj
H1gSOzLaGWbedZotMWU=
-----END-----
EOT
 }  }
postgresql_certs_by_ip = {  "10.1.0.101" = { private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
BOKnt5GgdKukzw2lBadXL/eF0yg8TNdmgRg1BsYZfFLWZyZ5FE2oqqKyOTWi2M5i
PbTvxWvIP4K8K8b1wnxLcFuN5OtJ21q+1MccprdfVOTMXg3d9ne2LDBkIPYshL+2
2RjhksDH7p7mbFap0jkUZJ+KKS/KVTEmkXpAW+zZVEWsmFoye2YhxPOw7uQRpkpw
bWMZsMyWMwKBgQDeMy65IzxdTaew4IScHoPMXY98F5D6l+uHW5erfYH//kH4o8qO
hgQ0oX3QuVS3cGdvC0yqVYuQwbxsObEYtwG8mO2pgpOUwGqPlrlKoYx6nBqDo4eq
YvXmWYr7l3Q+TOOdmSxVPv0EoE4BnEC/eL0cj4dx4JwG84fNulbYhgqTdQKBgQC1
P3PKJm6ElG3IIcsi6fLFsF6bq96Qxzzn5udfE4Z1bueNL+NBBq7fDTObMp22SM2V
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
dFNpxiedGkBWqtwJHYNu+kfbHETujMZAb80JYC3Q0MaBKKz1oAQ8yDu6YXHCjgIa
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
 }
  "10.1.1.237" = { private_key = <<-EOT
-----START-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC9l8i69P2WE00y
PD4zbWX+qEib095Snjrjs4PA6posCPpdRkI5zlc5ApOb5uFc7mJnomMtqJ1TQigQ
94iOC5Gyu+9NOpeLWrRYjMfhX2v3f/yskMhZ+Pdn/wLSZ2VfpAF81c9I8KhQZNpW
IsFffvZQiLUCgYA0C2uf21XaXGb4LHkXvBnD0MbVORNfQkMB03QQxns4ttwxKbLj
MQraLF8BG+RppMe7t8pIiqp7yCVU+cgoBFGNgGpG/8zI6XGEGuA5azffaAYEWfOK
DVK4Lmw+RSU+UBgPekN92UvPDlAuQtMkw78IPTvkOIf6mdhjeD9cPh0hEQKBgDuG
GgO2IQRw6z76GIzsJQbKXnGMjw4cgcx5Lcj7TbE9bvis5oivvi0j6uIH55iHaWpJ
JDY5yNo/zqrBEs4/0uUXiE6bQLV+hoGM4DB/D8rQxtsEjTWWlpe0gCtV/2MoSCiU
MPDvEqUsuMH1TtrZKvYM0TVaYZlko3kJKusze0iBAoGAFAU4vuLP7oXs21lVbQTX
ZQ+R3WD3Rff4Kjvf7HjseyQf/tKD+k3+mRNWvrna+xvw8O6tjhDjwiSS0bMymGHi
qvSDQj0+NynYBvrWqaCKujyC16psh+4rMisRjzHMa66pMq+BXju28beAEteoayG6
VIu99GxBWQY3QzBiFRm2hMU=
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzpMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAICL
j41rGqCTn3el+jnWWU40/YasNv5VdhFzHr7/t6YZQWb9MWV+22QcdqHoG+zt3zWO
Ha42TtaWNb42PWK107i3naXQR9XC3blfsawMqMYBP6Gt/p/b12WvMR9ZGdYYVhZl
ser9AsWiwAz7hSHEI6ih4hFsUAuyMdBX+y6lGcIaTj3XI/12R7EbXZl6+JKgSY5c
3h/Wv2VET772sWjwqxWHbzC9NEe4GWtw9j1QMOFniz79D0GiCrW1HzHK1ZaXioe7
LhZIOS8c5KyyrC9jrl88w38cU+nFQOXWvsecIm+HYba06v1515J4hPSADqsruiMf
yyGlcPRD4hI+ENvAWU8=
-----END-----
EOT
 }
  "10.1.2.21" = { private_key = <<-EOT
-----START-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC1yT32Ii+EUIAe
FZQbOxzGduI4I6NDicF0VTa22hW++Fdn7lrgq/CirwjlQNKXsImG+9zlQHVIuna3
eixWmWYgICUZ/kLYtdtuDiAvRpkj3zy1WeApgPgyuLUSb+woscCPn/FOuUH+I9I3
rAI2VYXU7P9j3T+LzMy5xZqISiPi5AAo81MY9OJoeR+yDnqZRfw5O4nDGsTERG69
ZfXa4S8jc9K+VjvOQdS/Ihj4hLovj5ROcYUltsTQQLbqgi4NL5jynNZegCP1mmOb
6hzksNiElwKBgQCO7xG9GKaoVXAiO+WK0Lybu19xi3H90lVo7jJ1MpAoLAuw2LoY
KM9RFhUht74fGPwTa+SGKp/HumrVltze9sge58hHDZCeo2y91Ui7oJHvGo8obD88
blXEu7TsLqAb08eAA1PEWhLmasCNUq8SPP8xfndviTFwvNUTneE4WkJ/uQKBgF6V
AUuNLWgBclV6TzYlm10Wo/FHcBSHyl1SOcgkTAdeCKRIm1Jd+T+LIUFcvtPOPMNG
4UAvlM496ruqJnhAAkXuHmdN1KW7zbNw/FW7PnHPQM7RTzkdgHmh3PQHRojZnTN6
ABNejk5jj31GwRc7Qm5JYALhq/GNs8YPsmIATiD7AoGAJQ1N36pQgYy5Jy5I0aIq
p3nGfRWBLYqftL2oyDQGeyRQX9/hF8hh6AwbET/CAPoV+e4EjOI9VeJVjWYZ1raB
CHBWACozIc+2fQQngj5PsH6PwP+sq47BcA8mEorJcYRelKZF1NJp9KSrcaDSo8MJ
w/qTq+SCthUyjrVT50EN34k=
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzqMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAH1a
5ER8bgk1E3600ak0y01E9Li3GbECwshLlx9OH4NesBuFYFuzI5gYd+pKQVnFsRiP
X8CJ0GLWrbfuOk27mIKkKFsuT44WQQMXWb4uf0h7oRENG8ASeiZHw9BmMUCPpufh
6WA1WxtHm4Bq+l3910YMaNZFZi953BUXF6OWfrMYelqlGEDR4UIkFGw5cmHBQJbN
Uw4+ATYf15lE3CQsKDknuRSJYun7M4+Wc8JKArtTPgVJ/la3Y1tpaKf36UL3qkxK
d0XhuTq+G0/vyjRCjAESDkF2Djw2Eb7VFnDPQCfRkKjjgIWjps0/kBVP6T4flQGj
H1gSOzLaGWbedZotMWU=
-----END-----
EOT
 }  }
opensearch_certs_by_ip = {  "10.1.0.221" = { private_key = <<-EOT
-----START-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzanmNSWM8/rEe
ev2e1SttSsDeb2ujywHhkk1EzT2I2vHJ3YW56MAPgTGbLEjD1UP/BUIbX1EY76No
BOKnt5GgdKukzw2lBadXL/eF0yg8TNdmgRg1BsYZfFLWZyZ5FE2oqqKyOTWi2M5i
PbTvxWvIP4K8K8b1wnxLcFuN5OtJ21q+1MccprdfVOTMXg3d9ne2LDBkIPYshL+2
b7NPQ3DPuS6FL6zwlLOJmyi7/AGFchAT/Z1R/ALdhm9CUHw04agsgCrYxFllb958
+2xSd/pUy1RRIYZgpop3TEsndg6Xklf5uOF9q4BmpiZpGGS/A/OtY8yHD/CeWQnh
FTOPHQKHpLehCms3ST5r3FOjnQKBgQDH7TfkMTncu4+3Z4TTFOeaMG5+ZRfHSZMo
RMyJRaCbwi6sOhV2JUdi3m3XIM+MNXrIl5XPWzOk9p2hMjIE+VqpwKNBT7gfucxu
2RjhksDH7p7mbFap0jkUZJ+KKS/KVTEmkXpAW+zZVEWsmFoye2YhxPOw7uQRpkpw
bWMZsMyWMwKBgQDeMy65IzxdTaew4IScHoPMXY98F5D6l+uHW5erfYH//kH4o8qO
hgQ0oX3QuVS3cGdvC0yqVYuQwbxsObEYtwG8mO2pgpOUwGqPlrlKoYx6nBqDo4eq
YvXmWYr7l3Q+TOOdmSxVPv0EoE4BnEC/eL0cj4dx4JwG84fNulbYhgqTdQKBgQC1
P3PKJm6ElG3IIcsi6fLFsF6bq96Qxzzn5udfE4Z1bueNL+NBBq7fDTObMp22SM2V
iQdBwzbguMtS1gTOyZjOCFx422LbMEALxjluVvL6znO0BkfTEUi6GjcgtBIlZnJ2
urCQJEZnCKVZwY0T8xc2Si17HDgbYBxgLGsJYWU4TQKBgQDUUr0wMccyAVdcWDSu
INgIogAUUe55Rb+PUZEg4GJfaBYj46W3pquJ1i760DBXWWybrbNP2i/mGhSOxpND
h5icuS3RCp1/M6zyvC7sd2RJOrneD6dpzXwofkxTcXEZJ4QicdpJGJxu/q0JeGd+
1oJBPS3l70y308n9UeYTsMjaYQ==
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzoMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
MjIxMTAzMDU0ODMyWhcNMjUxMTAyMDU0ODMyWjBjMQswCQYDVQQGEwJVUzETMBEG
l5JX+bjhfauAZqYmaRhkvwPzrWPMhw/wnlkJ4WcNN6OWRwIDAQABoyEwHzAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBALB9
dFNpxiedGkBWqtwJHYNu+kfbHETujMZAb80JYC3Q0MaBKKz1oAQ8yDu6YXHCjgIa
sma0aIEmqkebKEu1zu90EJc2tPBKyczEkDhYNq3qDIEJnpurw3/UneLMU78N8Q4G
G++oW+G/tdC3O1YpEOez4WFyolr5me6X/4PuFyJKIzEurcnyVt9Nw8OTNfT7FCLz
+/NFCuciSxtut6LzCWptIv38pPoSxodESP7ZbzoKM0Nh/APJNxa7TIo311bkTJz/
OFfsG3Ck7YUEQKhWvwmnRkF6Y22HLGnastXKArUUwhkoPCKX8leob8vtPPEaCNCS
I2/48tM1RZehjlEb8vM=
-----END-----
EOT
 nodes_dn = "CN=chefnode,O=Chef Software Inc,L=Seattle,ST=Washington,C=US" }
  "10.1.1.38" = { private_key = <<-EOT
-----START-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC9l8i69P2WE00y
PD4zbWX+qEib095Snjrjs4PA6posCPpdRkI5zlc5ApOb5uFc7mJnomMtqJ1TQigQ
94iOC5Gyu+9NOpeLWrRYjMfhX2v3f/yskMhZ+Pdn/wLSZ2VfpAF81c9I8KhQZNpW
s9Rr0OkN978RQLVie7WCbtPyuIUFBjkORbxJOE8HzjtCv4qjvmKZAXAe2FBPrDzU
XK7NWJdXSrubLRjoCXjt72bfCS8DD0b6n315UG/VKTQL/mProYIlO3MwO4tQBGpi
THDYwQBNneS2TcYqM+N7wudKV1UCgYEAwc9vPAibZzKqu7FXl95ym0LQVGNQhGum
KxPep2QAkQgkgPcJOnvQV4fVv4bWffPQm00bQrjDnM2lX9gVW8+VBdSfNpUnsQh+
+8UCAGXAomX5OyjN3XQmCEfvT6VgXu1sEXRRdtfAxCnk6wPgdblBuEOnTXn7xkk4
IsFffvZQiLUCgYA0C2uf21XaXGb4LHkXvBnD0MbVORNfQkMB03QQxns4ttwxKbLj
MQraLF8BG+RppMe7t8pIiqp7yCVU+cgoBFGNgGpG/8zI6XGEGuA5azffaAYEWfOK
DVK4Lmw+RSU+UBgPekN92UvPDlAuQtMkw78IPTvkOIf6mdhjeD9cPh0hEQKBgDuG
GgO2IQRw6z76GIzsJQbKXnGMjw4cgcx5Lcj7TbE9bvis5oivvi0j6uIH55iHaWpJ
JDY5yNo/zqrBEs4/0uUXiE6bQLV+hoGM4DB/D8rQxtsEjTWWlpe0gCtV/2MoSCiU
MPDvEqUsuMH1TtrZKvYM0TVaYZlko3kJKusze0iBAoGAFAU4vuLP7oXs21lVbQTX
ZQ+R3WD3Rff4Kjvf7HjseyQf/tKD+k3+mRNWvrna+xvw8O6tjhDjwiSS0bMymGHi
qvSDQj0+NynYBvrWqaCKujyC16psh+4rMisRjzHMa66pMq+BXju28beAEteoayG6
VIu99GxBWQY3QzBiFRm2hMU=
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzpMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
MjIxMTAzMDU0ODMyWhcNMjUxMTAyMDU0ODMyWjBjMQswCQYDVQQGEwJVUzETMBEG
BQY5DkW8SThPB847Qr+Ko75imQFwHthQT6w81FyuzViXV0q7my0Y6Al47e9m3wkv
Aw9G+p99eVBv1Sk0C/5j66GCJTtzMDuLUARqYlebH6D6HTIiyEuuWlPBdBfRfFbr
EuiI52kRrmj/W8ElRO0HeC86GF2lI4TFK2sO1Emp/oTnGQIDAQABoyEwHzAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAICL
j41rGqCTn3el+jnWWU40/YasNv5VdhFzHr7/t6YZQWb9MWV+22QcdqHoG+zt3zWO
Ha42TtaWNb42PWK107i3naXQR9XC3blfsawMqMYBP6Gt/p/b12WvMR9ZGdYYVhZl
ser9AsWiwAz7hSHEI6ih4hFsUAuyMdBX+y6lGcIaTj3XI/12R7EbXZl6+JKgSY5c
3h/Wv2VET772sWjwqxWHbzC9NEe4GWtw9j1QMOFniz79D0GiCrW1HzHK1ZaXioe7
LhZIOS8c5KyyrC9jrl88w38cU+nFQOXWvsecIm+HYba06v1515J4hPSADqsruiMf
yyGlcPRD4hI+ENvAWU8=
-----END-----
EOT
 nodes_dn = "CN=chefnode,O=Chef Software Inc,L=Seattle,ST=Washington,C=US" }
  "10.1.2.245" = { private_key = <<-EOT
-----START-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC1yT32Ii+EUIAe
FZQbOxzGduI4I6NDicF0VTa22hW++Fdn7lrgq/CirwjlQNKXsImG+9zlQHVIuna3
eixWmWYgICUZ/kLYtdtuDiAvRpkj3zy1WeApgPgyuLUSb+woscCPn/FOuUH+I9I3
rAI2VYXU7P9j3T+LzMy5xZqISiPi5AAo81MY9OJoeR+yDnqZRfw5O4nDGsTERG69
E7eSvSu/hlP9B7RNnzJ1py0Vc+BcKm4qjgR+g+kARwY6loMWyYjB0asMW/akbNTU
ls//BuB/lokXniZgJK7PbipptYzUgXjC8YRXA7SBRFKVRAKL6GTXwklSZxCBJCe+
MRkkGxffjK+G5qBeFUyy13KOVA6C8c+r+zRYtuWm/T9XpgIGWvqbxY1JNR11dBbx
T0254P2oDlCSOeNTt2kHJw1WvVScEVr9xOuCizs+JQKBgQDtDgYJ916DWz6hhPYD
0vVg5AINmPBUnB/Kiay8X0r8UlOTtcA04ek1VtMkcKvMojAFO/1Rfhujm0PDlKFS
ZfXa4S8jc9K+VjvOQdS/Ihj4hLovj5ROcYUltsTQQLbqgi4NL5jynNZegCP1mmOb
6hzksNiElwKBgQCO7xG9GKaoVXAiO+WK0Lybu19xi3H90lVo7jJ1MpAoLAuw2LoY
KM9RFhUht74fGPwTa+SGKp/HumrVltze9sge58hHDZCeo2y91Ui7oJHvGo8obD88
blXEu7TsLqAb08eAA1PEWhLmasCNUq8SPP8xfndviTFwvNUTneE4WkJ/uQKBgF6V
AUuNLWgBclV6TzYlm10Wo/FHcBSHyl1SOcgkTAdeCKRIm1Jd+T+LIUFcvtPOPMNG
4UAvlM496ruqJnhAAkXuHmdN1KW7zbNw/FW7PnHPQM7RTzkdgHmh3PQHRojZnTN6
ABNejk5jj31GwRc7Qm5JYALhq/GNs8YPsmIATiD7AoGAJQ1N36pQgYy5Jy5I0aIq
p3nGfRWBLYqftL2oyDQGeyRQX9/hF8hh6AwbET/CAPoV+e4EjOI9VeJVjWYZ1raB
CHBWACozIc+2fQQngj5PsH6PwP+sq47BcA8mEorJcYRelKZF1NJp9KSrcaDSo8MJ
w/qTq+SCthUyjrVT50EN34k=
-----END-----
EOT
 public_key = <<-EOT
-----START-----
MIIDajCCAlKgAwIBAgIJAPHEYb1O7wzqMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow
GAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN
MjIxMTAzMDU0ODMzWhcNMjUxMTAyMDU0ODMzWjBjMQswCQYDVQQGEwJVUzETMBEG
A1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl
I988tVngKYD4Mri1Em/sKLHAj5/xTrlB/iPSN6wCNlWF1Oz/Y90/i8zMucWaiEoj
4uQAKPNTGPTiaHkfsg56mUX8OTuJwxrExERuvXitLPMmZSRKKcikXicp5yqlFYaA
MdSFR+1VTDNwclVMvpmzoxhOby8r4P5YLfzEVDJ4wK1Qpki7vjtjnZiipvrWmJJM
QIzkWVvgj5JgtfHpmZ8XPBYLfl5ZcL5AxtM24TCvhwmghQIDAQABoyEwHzAdBgNV
HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDQYJKoZIhvcNAQELBQADggEBAH1a
5ER8bgk1E3600ak0y01E9Li3GbECwshLlx9OH4NesBuFYFuzI5gYd+pKQVnFsRiP
X8CJ0GLWrbfuOk27mIKkKFsuT44WQQMXWb4uf0h7oRENG8ASeiZHw9BmMUCPpufh
6WA1WxtHm4Bq+l3910YMaNZFZi953BUXF6OWfrMYelqlGEDR4UIkFGw5cmHBQJbN
Uw4+ATYf15lE3CQsKDknuRSJYun7M4+Wc8JKArtTPgVJ/la3Y1tpaKf36UL3qkxK
d0XhuTq+G0/vyjRCjAESDkF2Djw2Eb7VFnDPQCfRkKjjgIWjps0/kBVP6T4flQGj
H1gSOzLaGWbedZotMWU=
-----END-----
EOT
 nodes_dn = "CN=testnode,O=Chef Software Inc,L=Bangalore,ST=India,C=IN" }  }
nfs_mount_path = "/mnt/automate_backups"
postgresql_instance_count = 3
postgresql_archive_disk_fs_path = "/mnt/automate_backups/postgresql"
habitat_uid_gid = ""
ssh_user = "ec2-user"
ssh_port = "22"
ssh_key_file = "/home/ec2-user/A2HA.pem"
# ObjectStorage
################################################################################
bucket_name = "a2ha-test-bucket"
access_key = "test-access-key"
secret_key = "test-secret-key"
endpoint = "s3.amazonaws.com"
region = "ap-south-1"
