# flood-fill

modify image with some function. supports only jpeg and png images.

NOTE: i don't know if it's actually flood fill algorithm but name is cool tho :)

## usage

### convert dominant color with complementary one

```bash
go run . -n 1 -op c ./image.jpg
```

### convert every pixel with black & white

```bash
go run . -n 0 -op bw ./image.png
```

## docs

### number of dominant colors

```bash
go run . -n {number}
```

0 - means every color

### operation

```bash
go run . -op {bw|c}
```

bw - black & white
c - complementary
