# TypeID Specification (Version 0.3.0)

## Overview

TypeIDs are a type-safe extension of UUIDv7, they encode UUIDs in base32 and add
a type prefix.

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
   at most 63 characters in all lowercase snake_case ASCII `[a-z_]`.
1. A **separator**: an underscore `_` character. The separator is omitted if the
   prefix is empty.
1. A **UUID suffix**: a 128-bit UUIDv7 encoded as a 26-character string in
   base32.

### Length Constraints

The overall length of a TypeID is constrained by its components:

- **Minimum length**: 26 characters (when there is no prefix, just the UUID
  suffix)
- **Maximum length**: 90 characters (63 for prefix + 1 for separator + 26 for
  suffix)

### Type Prefix

A type prefix is a string denoting the type of the ID. The prefix must:

- Contain at most 63 characters.
- May be empty.
- If not empty:
  - Must contain only lowercase alphabetic ASCII characters `[a-z]`, or an
    underscore `_`.
  - Must start and end with an alphabetic character `[a-z]`. Underscores are not
    allowed at the beginning or end of the string.
  - **Digits and uppercase letters are not permitted.**
  - **Consecutive underscores (e.g., `foo__bar`) are allowed.**

Valid prefixes match the following regex: `^([a-z]([a-z_]{0,61}[a-z])?)?$`.

The empty string is a valid prefix, it's there for use cases in which
applications need to encode a typeid but elide the type information.

While short prefixes are technically allowed, for clarity and future
scalability, applications SHOULD use a prefix that is at least 3 characters long
(e.g., `usr_...` or `user_...` rather than `u_...`).

### Separator

The separator is a single underscore character `_`. If the prefix is empty, the
separator is omitted.

### UUID Suffix

The UUID suffix encodes exactly 128-bits of data in 26 characters. It uses the
base32 encoding described below.

#### Base32 Encoding

The UUID's 128 bits are treated in big-endian order (most significant bit
first). Two zeroed bits are prepended to the left of these 128-bits, resulting
in 130-bits of data. The 130-bits are then split from left to right into 26
groups of 5 bits each. Each 5-bit group is encoded as a single character in the
base32 alphabet, resulting in a total of 26 characters.

In practice this is most often done by using bit-shifting and a lookup table.
See the
[reference implementation encoding](https://github.com/jetify-com/typeid-go/blob/main/base32/base32.go)
for an example.

Note that this is different from the standard base32 encoding which encodes in
groups of 5 bytes (40 bits) and appends any padding at the end of the data.

The encoding uses the following alphabet `0123456789abcdefghjkmnpqrstvwxyz` as
specified by the following table:

| Value | Symbol | Value | Symbol | Value | Symbol | Value | Symbol |
| ----- | ------ | ----- | ------ | ----- | ------ | ----- | ------ |
| 0     | 0      | 8     | 8      | 16    | g      | 24    | r      |
| 1     | 1      | 9     | 9      | 17    | h      | 25    | s      |
| 2     | 2      | 10    | a      | 18    | j      | 26    | t      |
| 3     | 3      | 11    | b      | 19    | k      | 27    | v      |
| 4     | 4      | 12    | c      | 20    | m      | 28    | w      |
| 5     | 5      | 13    | d      | 21    | n      | 29    | x      |
| 6     | 6      | 14    | e      | 22    | p      | 30    | y      |
| 7     | 7      | 15    | f      | 23    | q      | 31    | z      |

This is the same alphabet used by
[Crockford's base32 encoding](https://www.crockford.com/base32.html), but in our
case the alphabet encoding is strict: always in lowercase, no hyphens allowed,
and we never decode multiple ambiguous characters to the same value.

Technically speaking, 26 characters in base32 can encode 130 bits of data, but
UUIDs are 128 bits. We therefore need to prevent overflow errors by not allowing
typeids that would represent a value that exceeds 128 bits.

Since the leading 2 bits of the 130-bit value are always zero (due to the two
zeroed bits we prepend), the first 5-bit group (i.e., the first base32
character) must never exceed decimal 7 (`0b0111`). Therefore, the maximum
possible suffix for a typeid is `7zzzzzzzzzzzzzzzzzzzzzzzzz`, which corresponds
to the maximum 128-bit value of a UUID (`0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF`).

Implementations MUST reject any suffix where the first character is greater than
`7`, as this would represent a value that exceeds 128 bits. The simplest way to
validate this is by checking that the first character of the suffix is `7` or
less.

#### Compatibility with UUID

When generating a new TypeID, the generated UUID suffix MUST decode to a valid
UUIDv7. This means:

- Bits 48-51 of the UUID MUST be `0111` (indicating version 7)
- Bits 64-65 of the UUID MUST be `10` (indicating the UUID variant)

Implementations SHOULD allow encoding/decoding of other UUID variants when the
bits are provided by end users. This makes it possible for applications to
encode other UUID variants like UUIDv1 or UUIDv4 at their discretion.

## Examples

Here are additional examples of valid TypeIDs with different prefix patterns:

| Prefix                                                            | Example TypeID                                                                               | Notes                              |
| ----------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | ---------------------------------- |
| _(empty)_                                                         | `01h5fskfsk4fpeqwnsyz5hj55t`                                                                 | No prefix, no separator            |
| `user`                                                            | `user_01h5fskfsk4fpeqwnsyz5hj55t`                                                            | Common use case                    |
| `my_type`                                                         | `my_type_01h5fskfsk4fpeqwnsyz5hj55t`                                                         | Prefix with an underscore          |
| `a_b_c`                                                           | `a_b_c_01h5fskfsk4fpeqwnsyz5hj55t`                                                           | Multiple underscores               |
| `my__type`                                                        | `my__type_01h5fskfsk4fpeqwnsyz5hj55t`                                                        | Consecutive underscores            |
| `a`                                                               | `a_01h5fskfsk4fpeqwnsyz5hj55t`                                                               | Single-letter prefix (discouraged) |
| `abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk` | `abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijk_01h5fskfsk4fpeqwnsyz5hj55t` | Maximum length prefix (63 chars)   |

The following table shows examples of invalid TypeIDs that might appear valid at
first glance:

| Invalid TypeID Example               | Reason Invalid                                        |
| ------------------------------------ | ----------------------------------------------------- |
| `PREFIX_01h5fskfsk4fpeqwnsyz5hj55t`  | Prefix contains uppercase letters (must be lowercase) |
| `12345_01h5fskfsk4fpeqwnsyz5hj55t`   | Prefix contains numbers (only [a-z_] allowed)         |
| `_prefix_01h5fskfsk4fpeqwnsyz5hj55t` | Prefix starts with underscore (must start with [a-z]) |
| `prefix__01h5fskfsk4fpeqwnsyz5hj55t` | Prefix ends with underscore (must end with [a-z])     |
| `prefix_0123456789ABCDEFGHJKMNPQRS`  | Suffix contains uppercase letters (must be lowercase) |
| `prefix_8zzzzzzzzzzzzzzzzzzzzzzzzz`  | Suffix exceeds maximum value (first char must be ≤7)  |

For comprehensive testing, please refer to the full set of examples in the
[valid.yml](valid.yml) and [invalid.yml](invalid.yml) files. These files contain
numerous test cases to help validate implementations, including edge cases and
corner cases not shown above.

## Versioning

This spec uses semantic versioning: `MAJOR.MINOR.PATCH`. The version is
incremented when the spec changes in a way that is not backwards compatible.

Libraries that implement this spec should also use semantic versioning.

## Validating Implementations

To assist library authors in validating their implementations, we provide:

- A [reference implementation in Go](https://github.com/jetify-com/typeid-go)
  with extensive testing.
- A [valid.yml](valid.yml) file containing a list of valid typeids along with
  their corresponding decoded UUIDs. For convenience, we also provide a
  [valid.json](valid.json) file containing the same data in JSON format.\
  When implementing, ensure that:
  1. Encoding the hex UUID produces the expected base32 suffix
  2. Decoding the base32 suffix produces the original hex UUID
- An [invalid.yml](invalid.yml) file containing a list of strings that are
  invalid typeids and should fail to parse/decode. For convenience, we also
  provide a [invalid.json](invalid.json) file containing the same data in JSON
  format.\
  When implementing, ensure that your implementation rejects all invalid
  strings.
