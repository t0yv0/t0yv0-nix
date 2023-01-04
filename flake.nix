{
  description = "A flake for t0yv0-nix binary";

  inputs = {
    nixpkgs.url = github:NixOS/nixpkgs/nixos-22.11;
  };

  outputs =
    { self,
      nixpkgs,
    }:

    let
      ver = "0.0.1";

      package = { system }:
        let
          pkgs = import nixpkgs { system = system; };
        in pkgs.buildGo118Module rec {
          name = "t0yv0-nix-${ver}";
          version = "${ver}";
          subPackages = [ "cmd/t0yv0-nix" ];
          vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
          src = ./.;
        };
    in {
      packages.x86_64-linux.default = package {
        system = "x86_64-linux";
      };
      packages.x86_64-darwin.default = package {
        system = "x86_64-darwin";
      };
      packages.aarch64-darwin.default = package {
        system = "aarch64-darwin";
      };
    };
}
