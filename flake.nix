{
  description = "Simple tunneler in Go that exposes local ports to the internet";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    systems = [
      "x86_64-linux"
      "aarch64-linux"
      "x86_64-darwin"
      "aarch64-darwin"
    ];

    forAllSystems = function: nixpkgs.lib.genAttrs systems (system: function nixpkgs.legacyPackages.${system});
  in {
    packages = forAllSystems (pkgs: rec {
      default = trok;
      trok = pkgs.callPackage ./default.nix {};
    });

    nixosModules.default = ./module.nix;

    devShells = forAllSystems (pkgs: {
      default = pkgs.callPackage ./shell.nix {};
    });
  };
}
