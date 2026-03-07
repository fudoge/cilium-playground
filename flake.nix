{
  description = "Devshell for cilium study";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
  in {
    devShells."${system}".default = let
      pkgs = import nixpkgs {inherit system;};
    in
      pkgs.mkShell {
        name = "Cilium";

        packages = with pkgs; [
          go
          grpc
          buf
          protobuf
          protoc-gen-go
          protoc-gen-go-grpc

          cilium-cli
          kubernetes-helm
        ];

        shellHook = ''
          echo "🐝 Cilium DevShell initialized!"
        '';
      };
  };
}
