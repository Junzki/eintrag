# Der Eintrag

A lightweight auth service written in Go.

## Usage
1. Execute all SQL files in the `migrations/` directory
2. Copy `etc/config.example.json`, and write your own `config.json` 
   1. Generate a secret key and fill into the `secret_key` field
   2. Add your own database connection to `database_connection_string` (Note: Only **PostgreSQL** supported)
3. Execute the `eintrag` executable with following command
   ```bash
   $ eintrag --config /path/to/your/config.json
   ```
   
## Requirements
1. To build this project, you are supposed to have **Go 1.12** or later installed.
2. **PostgreSQL** 10 or later with **[pgcrypto]** installed is required.


## Features
- [x] Password Authentication
- [x] JWT Generation
- [ ] User Management
- [ ] WebAuthn Support
- [ ] HMAC-TOTP Support


## Cryptography
- JSON Web Tokens are signed with HS256 algorithm.
- Passwords are hashed with [pgcrypto] `crypt` function with `bf` algorithm.


## License
GPLv3


[pgcrypto]: https://www.postgresql.org/docs/current/pgcrypto.html
