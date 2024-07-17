# TypeID

### A type-safe, K-sortable, globally unique identifier inspired by Stripe IDs

![License: Apache 2.0](https://img.shields.io/github/license/jetify-com/typeid) [![Join Discord](https://img.shields.io/discord/903306922852245526?color=7389D8&label=discord&logo=discord&logoColor=ffffff)](https://discord.gg/jetify) [![Built with Devbox](https://www.jetify.com/img/devbox/shield_galaxy.svg)](https://www.jetify.com/devbox/)

## What is it?

TypeIDs are a modern, type-safe extension of UUIDv7. Inspired by a similar use of prefixes
in Stripe's APIs.

TypeIDs are canonically encoded as lowercase strings consisting of three parts:

1. A type prefix (at most 63 characters in all lowercase snake_case ASCII [a-z_]).
2. An underscore '\_' separator
3. A 128-bit UUIDv7 encoded as a 26-character string using a modified base32 encoding.

Here's an example of a TypeID of type `user`:

```
  user_2x4y6z8a0b1c2d3e4f5g6h7j8k
  └──┘ └────────────────────────┘
  type    uuid suffix (base32)
```

A [formal specification](./spec) defines the encoding in more detail.

## Online Converter
You can try converting UUID to TypeID and back using Jetify's TypeID Converter. Paste your TypeID string to convert to UUID or put your prefix and UUID in this format: `prefix:UUID` to convert to TypeID.

### [jetify.com/typeid](https://www.jetify.com/typeid)

## Benefits

- **Type-safe:** you can't accidentally use a `user` ID where a `post` ID is expected. When debugging, you can
  immediately understand what type of entity a TypeID refers to thanks to the type prefix.
- **Compatible with UUIDs:** TypeIDs are a superset of UUIDs. They are based on the upcoming [UUIDv7 standard](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7). If you decode the TypeID and remove the type information, you get a valid UUIDv7.
- **K-Sortable**: TypeIDs are K-sortable and can be used as the primary key in a database while ensuring good
  locality. Compare to entirely random global ids, like UUIDv4, that generally suffer from poor database locality.
- **Thoughtful encoding**: the base32 encoding is URL safe, case-insensitive, avoids ambiguous characters, can be
  selected for copy-pasting by double-clicking, and is a more compact encoding than the traditional hex encoding used by UUIDs (26 characters vs 36 characters).

## Implementations

Implementations should adhere to the formal [specification](./spec).

Latest spec version: v0.3.0

### Official Implementations by `jetify`

| Language                                              | Status        | Spec Version |
| ----------------------------------------------------- | ------------- | ------------ |
| [Go](https://github.com/jetify-com/typeid-go)         | ✓ Implemented | v0.3         |
| [SQL](https://github.com/jetify-com/typeid-sql)       | ✓ Implemented | v0.2         |
| [TypeScript](https://github.com/jetify-com/typeid-js) | ✓ Implemented | v0.3         |

### Community Provided Implementations

| Language                                                      | Author                                                                                    | Spec Version                                                               |
|---------------------------------------------------------------|-------------------------------------------------------------------------------------------|----------------------------------------------------------------------------|
| [C# (.Net)](https://github.com/TenCoKaciStromy/typeid-dotnet) | [@TenCoKaciStromy](https://github.com/TenCoKaciStromy)                                    | v0.2 on 2023-06-30                                                         |
| [C# (.Net Standard 2.1)](https://github.com/cbuctok/typeId)   | [@cbuctok](https://github.com/cbuctok)                                                    | v0.2 on 2023-07-03                                                         |
| [C# (.NET)](https://github.com/firenero/TypeId)               | [@firenero](https://github.com/firenero)                                                  | v0.3 on 2024-04-15                                                         |
| [Dart](https://github.com/TBD54566975/typeid-dart)            | [@mistermoe](https://github.com/mistermoe) [@tbd54566975](https://github.com/tbd54566975) | [v0.3 on 2024-07-02](https://github.com/TBD54566975/typeid-dart/actions/runs/9755701869/job/26924658060#step:6:10)                                                         |
| [Elixir](https://github.com/sloanelybutsurely/typeid-elixir)  | [@sloanelybutsurely](https://github.com/sloanelybutsurely)                                | v0.3 on 2024-04-22                                                         |
| [Elixir](https://github.com/xinz/elixir_typeid)               | [@xinz](https://github.com/xinz)                                                          | v0.1 on 2024-06-03                                                         |
| [Haskell](https://github.com/MMZK1526/mmzk-typeid)            | [@MMZK1526](https://github.com/MMZK1526)                                                  | v0.3 on 2024-04-19                                                         |
| [Java](https://github.com/fxlae/typeid-java)                  | [@fxlae](https://github.com/fxlae)                                                        | v0.3 on 2024-04-14                                                         |
| [Java](https://github.com/softprops/typeid-java)              | [@softprops](https://github.com/softprops)                                                | v0.2 on 2023-07-04                                                         |
| [Kotlin](https://github.com/aleris/typeid-kotlin)             | [@aleris](https://github.com/aleris)                                                      | v0.3 on 2024-05-18                                                         |
| [OCaml](https://github.com/titouancreach/typeid-ocaml)        | [@titouancreach](https://github.com/titouancreach)                                        | v0.3 on 2024-04-22                                                         |
| [PHP](https://github.com/BombenProdukt/typeid)                | [@BombenProdukt](https://github.com/BombenProdukt)                                        | v0.2 on 2023-07-03                                                         |
| [Postgres](https://github.com/blitss/typeid-postgres)         | [@blitss](https://github.com/blitss)                                                      | [v0.3 on 2024-06-24](https://github.com/blitss/typeid-postgres/actions/runs/9637303320/job/26576304134#step:11:288)                                                                                                                                                                            |
| [Python](https://github.com/akhundMurad/typeid-python)        | [@akhundMurad](https://github.com/akhundMurad)                                            | v0.2 on 2023-06-30                                                         |
| [Ruby](https://github.com/broothie/typeid-ruby)               | [@broothie](https://github.com/broothie)                                                  | [v0.3 on 2024-04-13](https://github.com/broothie/typeid-ruby/pull/17)      |
| [Rust](https://github.com/conradludgate/type-safe-id)         | [@conradludgate](https://github.com/conradludgate)                                        | [v0.3 on 2024-04-12](https://github.com/conradludgate/type-safe-id/pull/1) |
| [Rust](https://github.com/johnnynotsolucky/strong_id)         | [@johnnynotsolucky](https://github.com/johnnynotsolucky)                                  | [v0.3 on 2024-05-17](https://github.com/johnnynotsolucky/strong_id/commit/10aa50487bbdd851c58a2ed73071a50452441370) |
| [Scala](https://github.com/ant8e/uuid4cats-effect)            | [@ant8e](https://github.com/ant8e)                                                        | v0.3 on 2024-04-19                                                         |
| [Scala](https://github.com/guizmaii-opensource/zio-uuid)      | [@guizmaii](https://github.com/guizmaii)                                                  | Not validated yet                                                          |
| [Swift](https://github.com/Frizlab/swift-typeid)              | [@Frizlab](https://github.com/Frizlab)                                                    | v0.3 on 2024-04-19                                                         |
| [T-SQL](https://github.com/uniteeio/typeid_tsql)              | [@uniteeio](https://github.com/uniteeio)                                                  | v0.2 on 2023-08-25                                                         |
| [TypeScript](https://github.com/ongteckwu/typeid-ts)          | [@ongteckwu](https://github.com/ongteckwu)                                                | v0.2 on 2023-06-30                                                         |
| [Zig](https://github.com/tensorush/zig-typeid)                | [@tensorush](https://github.com/tensorush)                                                | v0.2 on 2023-07-05                                                         |

We are looking for community contributions to implement TypeIDs in other languages.

## Command-line Tool

This repo includes a command-line tool for generating TypeIDs. To install it, run:

```bash
curl -fsSL https://get.jetify.com/typeid | bash
```

To generate a new TypeID, run:

```console
$ typeid new prefix
prefix_01h2xcejqtf2nbrexx3vqjhp41
```

To decode an existing TypeID into a UUID run:

```console
$ typeid decode prefix_01h2xcejqtf2nbrexx3vqjhp41
type: prefix
uuid: 0188bac7-4afa-78aa-bc3b-bd1eef28d881
```

And to encode an existing UUID into a TypeID run:

```console
$ typeid encode prefix 0188bac7-4afa-78aa-bc3b-bd1eef28d881
prefix_01h2xcejqtf2nbrexx3vqjhp41
```

## Related Work

- [UUIDv7](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7) - The upcoming UUID standard that TypeIDs are based on.

Alternatives to UUIDv7 that are also worth considering (but not type-safe like TypeIDs):

- [xid](https://github.com/rs/xid)
- [ulid](https://github.com/ulid)
- [ksuid](https://github.com/segmentio/ksuid)
