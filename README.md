# gqq: goodreads quotes query

## motivation

didn't see any (simple) goodreads quotes scrapers out there so i tried making one.

## goals

- eventually make it UNIXesque
- scrape books quotes too not just author quotes
- glam it up with some [charmbraclet](https://github.com/charmbracelet/) packages

## jq magic (for the impatient)

> i use the fish shell so POSIX people you're going to have to adapt this to your shell of choice

- get the length of the quotes array

```fish
jq 'length' quotes.json
```

- get a random quote

```fish
echo (random 0 (jq 'length' quotes.json)) | xargs -I {} jq '.[{}]' quotes.json
```
