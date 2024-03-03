Traefik Certificate Manager
=========
[![CI](https://github.com/CastawayEGR/traefik-certificate-manager/actions/workflows/ci.yml/badge.svg)](https://github.com/CastawayEGR/traefik-certificate-manager/actions/workflows/ci.yml)
[![CD](https://github.com/CastawayEGR/traefik-certificate-manager/actions/workflows/cd.yml/badge.svg)](https://github.com/CastawayEGR/traefik-certificate-manager/actions/workflows/cd.yml)
[![MIT License](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![GitHub repo size in bytes](https://img.shields.io/github/repo-size/CastawayEGR/traefik-certificate-manager.svg?logoColor=brightgreen)](https://github.com/CastawayEGR/traefik-certificate-manager)
[![GitHub last commit](https://img.shields.io/github/last-commit/CastawayEGR/traefik-certificate-manager.svg?logoColor=brightgreen)](https://github.com/CastawayEGR/traefik-certificate-manager)

Traefik Certificate Manager is a command-line tool for managing the Let's Encrypt ACME JSON configuration file used by Traefik for SSL/TLS certificate management.

## Features

- Load an existing Traefik ACME JSON configuration file.
- View and delete domains and SANs (Subject Alternative Names) from the configuration file.
- Save the updated configuration file.


## Installation

`tcm` is available from the [project's releases page](https://github.com/castawayegr/traefik-certificate-manager/releases).

## Usage

```bash
tcm -f /path/to/acme.json
```

Once you run the command, you'll be prompted to select whether you want to work with domains or SANs. After selecting the option, you'll be presented with a list of available values. You can then select the values using the space bar for the values you want to remove from the configuration followed by enter to remove.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Author Information

This application was created by [Michael Tipton](https://ibeta.org).
