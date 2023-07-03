# TypeID Specification (Version 0.1.0)

## Overview
TypeIDs are a type-safe extension of UUIDv7, they encode UUIDs in base32 and add a type prefix.

Here's an example of a TypeID of type `user`:

```
  user_2x4y6z8a0b1c2d3e4f5g6h7j8k
  └──┘ └────────────────────────┘
  type    uuid suffix (base32)
```

This document formalizes the specification for TypeIDs.

## Specification

A typeid consists of three parts:
1. A **type prefix**: a string denoting the type of the ID. The prefix should be
   at most 63 characters in all lowercase ASCII [a-z].
1. A **separator**: an underscore '_' character.
1. A **UUID suffix**: a 128-bit UUIDv7 encoded as a 26-character string in base32.

### Type Prefix
A type prefix is a string denoting the type of the ID. The prefix should be at most
63 characters in all lowercase ASCII [a-z]. Valid prefixes should match the following
regex: `[a-z]{0,63}`.

The empty string is a valid prefix, it's there for very specific use cases in which
applications need to encode a typeid but elide the type information. In general though,
applications should use a prefix that is at least 3 characters long.

> Note: [There's a proposal](https://github.com/jetpack-io/typeid/issues/7) to add `_` as
> an allowed separator within type prefixes.

### Separator
The separator is a single underscore character `_`. If the prefix is empty, the separator
is omitted.

### UUID Suffix
The UUID suffix encodes exactly 128-bits of data in 26 characters. It uses the base32
encoding described below.

#### Base32 Encoding
Bytes from the UUID are encoded from left to right. Two zeroed bits are pre-pended
to the 128-bits of the UUID, resulting in 130-bits of data. The 130-bits are then
split into 5-bit chunks, and each chunk is encoded as a single character in the
base32 alphabet, resulting in a total of 26 characters.

In practice this is most often done by using bit-shifting and a lookup table. See
the [reference implementation encoding](https://github.com/jetpack-io/typeid-go/blob/main/base32/base32.go)
for an example.

Note that this is different from the standard base32 encoding which encodes in
groups of 5 bytes (40 bits) and appends any padding at the end of the data.

The encoding uses the following alphabet `0123456789abcdefghjkmnpqrstvwxyz` as
specified by the following table:

| Value | Symbol | Value | Symbol | Value | Symbol | Value | Symbol |
|-------|--------|-------|--------|-------|--------|-------|--------|
| 0     | 0      | 8     | 8      | 16    | g      | 24    | r      |
| 1     | 1      | 9     | 9      | 17    | h      | 25    | s      |
| 2     | 2      | 10    | a      | 18    | j      | 26    | t      |
| 3     | 3      | 11    | b      | 19    | k      | 27    | v      |
| 4     | 4      | 12    | c      | 20    | m      | 28    | w      |
| 5     | 5      | 13    | d      | 21    | n      | 29    | x      |
| 6     | 6      | 14    | e      | 22    | p      | 30    | y      |
| 7     | 7      | 15    | f      | 23    | q      | 31    | z      |

This is the same alphabet used by [Crockford's base32 encoding](https://www.crockford.com/base32.html),
but in our case the alphabet encoding is strict: always in lowercase, no hyphens allowed,
and we never decode multiple ambiguous characters to the same value.

#### Compatibility with UUID
When genarating a new TypeID, the generated UUID suffix MUST decode to a valid UUIDv7.

Implementations MAY allow encoding/decoding of other UUID variants when the
bits are provided by end users. This makes it possible for applications to encode
other UUID variants like UUIDv1 or UUIDv4 at their discretion.

## Versioning
This spec uses semantic versioning: `MAJOR.MINOR.PATCH`. The version is incremented
when the spec changes in a way that is not backwards compatible.

Libraries that implement this spec should also use semantic versioning, and their
MAJOR and MINOR versions should match the version of the spec they implement.
The PATCH version is up to the discretion of the library author.

## Validating Implementations
To assist library authors in validating their implementations, we provide:
+ A reference implementation in [Go](https://github.com/jetpack-io/typeid-go)
  with extensive testing.
+ A [valid.yml](valid.yml) file containing a list of valid typeids along 
  with their corresponding decoded UUIDs.
+ An [invalid.yml](invalid.yml) file containing a list of strings that are
  invalid typeids and should fail to parse/decode.
