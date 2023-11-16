# pfsf

PreFix SufFix a binary to add / remove prefix and/or suffix to an input string

## Install

```
go install github.com/shoxxdj/pfsf@latest
```

## Usage

```
PreFix SufFix a binary to add / remove prefix and/or suffix to an input string
        -i: input value (or Stdin)
        -p: Prefix value
        -s: Suffix value
---- ---- ---- ---- Examples ---- ---- ---- ----
|   pfsf -i website -p https:// -s .com        |
|   echo 'website' | pfsf -p https:// -s .com  |
---- ---- ---- ---- -------- ---- ---- ---- ----
```
