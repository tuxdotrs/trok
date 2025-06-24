<h3 align="center">
  trok
</h3>
<p align="center">
  Accessing your local service should be simple
</p>
<p align="center">
  <a href="https://wakatime.com/badge/user/012e8da9-99fe-4600-891b-bd9d8dce73d9/project/52396aaa-6648-4ee3-a470-7f02ce8d30b9"><img src="https://wakatime.com/badge/user/012e8da9-99fe-4600-891b-bd9d8dce73d9/project/52396aaa-6648-4ee3-a470-7f02ce8d30b9.svg" alt="wakatime"></a>
  <a href="https://builtwithnix.org" target="_blank"><img alt="home" src="https://img.shields.io/static/v1?logo=nixos&logoColor=white&label=&message=Built%20with%20Nix&color=41439a"></a>
  <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/tuxdotrs/nix-config">
  <img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/tuxdotrs/trok">
</p>

### Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Selfhost](#selfhost)

### Installation

```sh
curl -fsSL https://trok.cloud/install.sh | sh
```

#### Nix Flake

```nix
# Add to your flake inputs
trok = {
  url = "github:tuxdotrs/trok";
  inputs.nixpkgs.follows = "nixpkgs";
};

# Add this in your nixos config
environment.systemPackages = [ inputs.trok.packages.${system}.default ];
```

### Usage

```sh
trok tcp PORT_NUMBER
```

### Selfhost

TODO
