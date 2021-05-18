# vault-dotenv

Creates a .env file from vault secrets

## Config

Configuration can be done through env variables or programmatically through the `Config` object
The following env variables are supported:

```bash
VAULT_ADDR                          # Vault server URL (default "http://localhost:8200")
VAULT_CACERT                        # Path to CA file
VAULT_TOKEN                         # Vault Token
VAULT_ROLEID or VAULT_ROLE_ID       # Vault app role id
VAULT_SECRETID or VAULT_SECRET_ID   # Vault app role secret id
VAULT_MOUNTPOINT                    # Vault app role mountpoint (default "approle")
VAULT_CLIENT_TIMEOUT                # Client timeout
VAULT_SKIP_VERIFY                   # Do not check SSL
```
