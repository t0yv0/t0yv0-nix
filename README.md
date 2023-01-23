# t0yv0-nix

Wraps  [Nix](https://nixos.org) package manager CLI.

## List packages in the current profile

Equivalent to `nix profile list` but with a more readable output:

```
$ t0yv0-nix pl
0  t0yv0-emacs         0.0.1   github:t0yv0/t0yv0-emacs  86c9ff8b5f20
1  t0yv0-nix           0.0.2   ~/code/t0yv0-nix          e8b4455ab3c9
2  t0pu                0.0.2   ~/code/pulumi-nix-profile 20b00d1e9e18
3  password-store      1.7.4   ~/code/pulumi-nix-profile 20b00d1e9e18
4  mgitstatus          0.0.1   ~/code/pulumi-nix-profile 20b00d1e9e18
5  silver-searcher     2.2.0   ~/code/pulumi-nix-profile 20b00d1e9e18

```

## Upgrade all packages in the current profile

Shorthand to `nix profile upgrade '.*'`:

```
t0yv0-nix ua
```
