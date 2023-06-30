# TypeID

### A type-safe, K-sortable, globally unique identifier inspired by Stripe IDs

![License: Apache 2.0](https://img.shields.io/github/license/jetpack-io/typeid) [![Join Discord](https://img.shields.io/discord/903306922852245526?color=7389D8&label=discord&logo=discord&logoColor=ffffff)](https://discord.gg/agbskCJXk2) [![Built with Devbox](https://jetpack.io/img/devbox/shield_galaxy.svg)](https://jetpack.io/devbox/)

## What is it?
TypeIDs are a modern, type-safe extension of UUIDv7. Inspired by a similar use of prefixes
in Stripe's APIs.

TypeIDs are canonically encoded as lowercase strings consisting of three parts:
1. A type prefix (at most 63 characters in all lowercase ASCII [a-z])
2. An underscore '_' separator
3. A 128-bit UUIDv7 encoded as a 26-character string using a modified base32 encoding.

Here's an example of a TypeID of type `user`:

```
  user_2x4y6z8a0b1c2d3e4f5g6h7j8k
  └──┘ └────────────────────────┘
  type    uuid suffix (base32)
```

A [formal specification](./spec) defines the encoding in more detail.

## Benefits
+ **Type-safe:** you can't accidentally use a `user` ID where a `post` ID is expected. When debugging, you can
  immediately understand what type of entity a TypeID refers to thanks to the type prefix.
+ **Compatible with UUIDs:** TypeIDs are a superset of UUIDs. They are based on the upcoming [UUIDv7 standard](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7). If you decode the TypeID and remove the type information, you get a valid UUIDv7.
+ **K-Sortable**: TypeIDs are K-sortable and can be used as the primary key in a database while ensuring good
  locality. Compare to entirely random global ids, like UUIDv4, that generally suffer from poor database locality.
+ **Thoughtful encoding**: the base32 encoding is URL safe, case-insensitive, avoids ambiguous characters, can be
  selected for copy-pasting by double-clicking, and is a more compact encoding than the traditional hex encoding used by UUIDs (26 characters vs 36 characters).

## Implementations
Implementations should adhere to the formal [specification](./spec).

### Official Implementations by `jetpack.io`
| Language | Status |
| -------- | ------ |
| [Go](https://github.com/jetpack-io/typeid-go) | ✓ Implemented |
| Python | ... Coming Soon |
| [SQL](https://github.com/jetpack-io/typeid-sql) | ✓ Implemented |
| [TypeScript](https://github.com/jetpack-io/typeid-ts) | ✓ Implemented |

### Community Provided Implementations
| Language | Author | Validated Against Spec? |
| -------- | ------ | ---------------------- |
| [C# (.Net)](https://github.com/TenCoKaciStromy/typeid-dotnet) | @TenCoKaciStromy | Not Yet |
| [Python](https://github.com/akhundMurad/typeid-python) | @akhundMurad | Not Yet |
| [Rust](https://github.com/alisa101rs/typeid-rs) | @alisa101rs | Not Yet |
| [Rust](https://github.com/conrad/type-safe-id) | @conrad | Not Yet |
| [Swift](https://github.com/Frizlab/swift-typeid) | @Frizlab | Not Yet |
| [TypeScript](https://github.com/ongteckwu/typeid-ts) | @ongteckwu | Yes, on 2023-06-30 |

We are looking for community contributions to implement TypeIDs in other languages.

## Command-line Tool
This repo includes a command-line tool for generating TypeIDs. To install it, run:

```bash
curl -fsSL https://get.jetpack.io/typeid | bash
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
+ [UUIDv7](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7) - The upcoming UUID standard that TypeIDs are based on.

Alternatives to UUIDv7 that are also worth considering (but not type-safe like TypeIDs):
+ [xid](https://github.com/rs/xid)
+ [ulid](https://github.com/ulid)
+ [ksuid](https://github.com/segmentio/ksuid)
