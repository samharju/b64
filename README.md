# b64

There are builtin tools available for base64 encoding-usecases in all sane programming languages.
But why not write one by myself just for the kicks?

I read through https://en.wikipedia.org/wiki/Base64 and started typing.

## Install
```bash
# requires go, tried with 1.18+
go install github.com/samharju/b64@latest
```

## Usage

```bash
usage:
        b64 [-d] INPUTSTR
        b64 [-d] [-o <path>] INPUTSTR
        b64 [-d] [-i <path>]
        b64 [-d] [-i <path>] [-o <path>]

Reads from stdin and prints to stdout, unless input- or outputfile is given.
Encodes by default, decode with -d.

  -d    decode input
  -i string
        read from file instead of stdin
  -o string
        write to file instead of stdout
```

## Example
```bash
# create some poorly printable binary data
echo -n -e \\0xa9\\0xad\\0xff\\0xa9\\0xf1 > binfile
cat binfile
# encode and transmit over some medium as string (logs etc):
body="$(b64 -i binfile)"
echo "\nencoded: $body"
# verify decoded content to be the same as source with hexdump
echo "original:"
hexdump -x binfile
echo "after encode/decode:"
echo -n -e $(b64 -d "$body") | hexdump -x

# �����
# encoded: qa3/qfE=
# original:
# 0000000    ada9    a9ff    00f1
# 0000005
# after encode/decode:
# 0000000    ada9    a9ff    00f1
# 0000005
```
