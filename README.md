# Der Eintrag

A lightweight auth service written in Go.

## Usage
1. Execute all SQL files in the `migrations/` directory.
2. Generate configuration JSON with config generator. This will create a new `config.json` at your preferred location, with unique random secret key attached.
   ```bash
   $ eintrag-config --config /path/to/your/config.json
   ```
   
3. Add your own database connection to `database_connection_string` to your `config.json` (Note: Only **PostgreSQL** supported)
4. Execute the `eintrag` executable with following command
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
