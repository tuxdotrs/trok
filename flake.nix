{
  description =
    "Simple tunneler in Go that exposes local ports to the internet";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [ inputs.treefmt-nix.flakeModule ];

      systems =
        [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];

      perSystem = { config, pkgs, ... }: {
        treefmt = {
          projectRootFile = "flake.nix";

          programs = {
            nixfmt.enable = true;
            gofumpt.enable = true;
          };
        };

        packages = rec {
          default = tfolio;
          tfolio = pkgs.callPackage ./nix/package.nix { };
        };

        devShells.default = pkgs.callPackage ./nix/shell.nix { };
      };

      flake.nixosModules.default = ./nix/module.nix;
    };

}
