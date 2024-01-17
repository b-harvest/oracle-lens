# Oracle Lens

Oracle Lens is monitoring program for operator of program which bring off-chain data to on-chain such as Oracle, Bridge, ETC.

## Monitoring Support Chains

- Injective
  - Peggo
- Umee
  - Oracle

## Quick Guide

1. Build

```bash
go build
```

2. Configure config.toml file

```bash
# You can get a example of config.toml file by below command
cp config.toml.example config.toml
```

3. Execute

```bash
./oracle-lens -config config.toml
```
