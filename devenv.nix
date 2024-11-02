{ pkgs, lib, config, inputs, ... }:

{
  packages = with pkgs; [ gopls nodejs tinygo go godef ];
}
