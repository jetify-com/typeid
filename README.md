# TypeID

### A type-safe, K-sortable, globally unique identifier inspired by Stripe IDs

![License: Apache 2.0](https://img.shields.io/github/license/jetpack-io/typeid)

## What is it?
TypeIDs are a modern, type-safe extension of UUIDv7.

TypeIDs are canonically encoded as lowercase strings consisting of three parts:
1. A type prefix
2. An underscore '_' separator
3. A 128-bit UUIDv7 encoded as a 26-character string in base32.

Here's an example of a TypeID of type `user`:

```
  user_2x4y6z8a0b1c2d3e4f5g6h7j8k
  └──┘ └────────────────────────┘
  type       uuid (base32)
```

## Benefits
+ Type-safe: you can't accidentally use a `user` ID where a `post` ID is expected. When debugging, you can
  immediately understand what type of entity a TypeID refers to thanks to the type prefix.
+ Standards-based: they are based on the upcoming [UUIDv7 standard](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7). If you remove the type information, you get a valid UUIDv7.
+ They are K-sortable and can be used as the primary key in a database. Entirely random global ids generally
  suffer from poor database locality.
+ Thoughtful encoding: the base32 encoding is URL safe, case-insensitive, avoids ambiguous characters, can be
  selected for copy-pasting by double-clicking, and is a more compact encoding than the traditional hex encoding used by UUIDs (26 characters vs 36 characters).


## Related Work
+ [UUIDv7](https://www.ietf.org/archive/id/draft-peabody-dispatch-new-uuid-format-04.html#name-uuid-version-7) - The upcoming UUID standard that TypeIDs are based on.

Alternatives to UUIDv7 that are also worth considering (but not type-safe like TypeIDs):
+ [xid](https://github.com/rs/xid)
+ [ulid](https://github.com/ulid)
+ [ksuid](https://github.com/segmentio/ksuid)