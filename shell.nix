{
  callPackage,
  go,
}: let
  mainPkg = callPackage ./default.nix {};
in
  mainPkg.overrideAttrs (oa: {
    nativeBuildInputs =
      [
        go
      ]
      ++ (oa.nativeBuildInputs or []);
  })
