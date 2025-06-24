{
  lib,
  buildGoModule,
}:
buildGoModule {
  pname = "trok";
  version = "0.2.0";

  src = ./.;
  vendorHash = "sha256-KwcLxkW3pbzujjc6JOZRwATlAA/qndf4FWpkJANv2z8=";

  meta = {
    license = lib.licenses.mit;
    maintainers = with lib.maintainers; [tuxdotrs];
    mainProgram = "trok";
  };
}
