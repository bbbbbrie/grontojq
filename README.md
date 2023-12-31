# ✨ grontojq

## 📦 Installation

```
go install github.com/bbbbbrie/grontojq@latest
```

That should put the `grontojq` executable in `$GOBIN`. If that's in `$PATH`, `which grontojq` should point you to it. 

## ℹ️  How one might use `grontojq`

This is a Go tool that converts `gron` output into `jq` commands. I wrote this to solve a specific problem I had and as an excuse to write in Go again. 

I typically use `gron` to help build the `jq` command I want to use for whatever I'm trying to do. While `--ungron` is nice, I don't want to have `gron` as a dependency in pipelines or other places where I'm making  use of `jq`.  I use `gron` to help build the `jq` commands. Once I have them, I use just `jq` for the actual parsing. 

You can pipe the output of `gron` in to `grontojq`. Here's how I think `grontojq` is most useful:

**First**, you'll use `gron cats.json` to explore the  JSON file you're interested in:

```
# gron cats.json
...
json.catDescriptions[10] = {};
json.catDescriptions[10].emoji = "😽";
json.catDescriptions[10].word = "Charming";
json.catDescriptions[11] = {};
json.catDescriptions[11].emoji = "😈";
json.catDescriptions[11].word = "Mischievous";
```

**Next**, you'll do something like `gron cats.json | grep 11 | grep word` to pick the exact line you're interested in:

```
# gron cats.json | grep 11 | grep word
json.catDescriptions[11].word = "Mischievous";
```

**Then**, you'll pass that output to `grontojq`:

```
# gron cats.json | grep 11 | grep word | grontojq
jq '.catDescriptions[11].word'
```

You can copy this from your terminal or use `pbcopy` or similar to copy it straight to your clipboard with something like `gron cats.json | grep 11 | grep word | grontojq | pbcopy`.

**Finally**, you'll take the output from `grontojq` and use it to parse the file more elegantly:

```
jq '.catDescriptions[11].word' cats.json
```

## 🐈‍⬛ Why?

This makes it a tiny bit easier  to go from `gron` output to `jq` filter -- especially if piping to something like `grontojq` is easy to add to your terminal muscle memory. I'd alias it to something like `g2jq` or similar to make  that even easier.