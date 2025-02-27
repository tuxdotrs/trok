{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "trok";
  version = "0.1.0";

  src = ./.;
  vendorHash = "sha256-P7UBLLMMOwCeiM7aAY6A8cyXqxwn7mDTNS2kJTALHPU=";

  meta = {
    license = lib.licenses.mit;
    maintainers = with lib.maintainers; [tuxdotrs];
    mainProgram = "trok";
  };
}
