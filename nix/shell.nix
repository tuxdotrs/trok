{ callPackage, go, }:
let mainPkg = callPackage ./package.nix { };
in mainPkg.overrideAttrs
(oa: { nativeBuildInputs = [ go ] ++ (oa.nativeBuildInputs or [ ]); })
