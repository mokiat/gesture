# gesture

A gesture recognition library written in [Go](https://go.dev/) that is based on the `$1`, `$P`, and `protractor` algorithm papers, with some small modifications.

For the curious, make sure to check the following papers (especially the protractor one):

- [$1](https://faculty.washington.edu/wobbrock/pubs/uist-07.01.pdf) - explains the $1 algorithm ([web page](https://depts.washington.edu/acelab/proj/dollar/index.html))
- [$P](https://faculty.washington.edu/wobbrock/pubs/icmi-12.pdf) - explains the $P algorithm ([web page](https://depts.washington.edu/acelab/proj/dollar/pdollar.html))
- [protractor](https://dl.acm.org/doi/10.1145/1753326.1753654) - explains the protractor algorithm

## User's Guide

There are two main concepts in this library - `Path` and `Pattern`.

The `Path` type represents a sequence of points in 2D space. These could come from a user's mouse gesture or could be specified manually. It can have an arbitrary length and as such has dynamic memory footprint.

The `Pattern` type represents a gesture pattern that can be matched for similarity against other patterns. It has a fixed memory footprint and is designed to be performant to use.

One would create a `Pattern` using a `Path` as follows:

```go
patternSquare := gesture.CreatePattern(gesture.Path{
  dprec.NewVec2(0.0, 0.0), // lower-left
  dprec.NewVec2(2.0, 0.0), // lower-right
  dprec.NewVec2(2.0, 2.0), // upper-right
  dprec.NewVec2(0.0, 2.0), // upper-left
  dprec.NewVec2(0.0, 0.0), // lower-left (in order to close the path)
})
```

Similarly, patterns for circles, triangles and any custom shapes can be defined this way.

Alternatively, one would accumulate points into a `Path` and create a query `Pattern` as follows:

```go
var path gesture.Path

someUserInterface.OnMouseMove(func(x, y int) {
  path = append(path, dprec.NewVec2(float64(x), float64(y)))
  // You could use some techniques here to reduce allocation by performing
  // path resampling. Check the Path API in the gesture package.
})

someUserInterface.OnMouseClick(func(x, y int) {
  if len(path) > 0 {
    pattern := gesture.CreatePattern(path)

    // Normally you would want to check against other shapes as well.
    // You might even have a dummy "no-op" shape to catch non-shapes.
    similarity := gesture.SimilarityDot(patternSquare, pattern)
    if similarity > 0.9 {
      fmt.Println("Might be a square")
    }
  }
  path = path[:0] // reset the path and reuse in future gestures
})
```

## Limitations

- **The implementation uses a hardcoded sampling resolution of 64** - the `protractor` paper mentions 32 as a good pick. Here 64 is used to be on the safe side.
- **The library uses float64** - This could have some performance implications compared to 32bit though most modern machines deal fairly well with 64bit floats.
- **The implementation is orientation invariant** - rotating a gesture will still match it.
- **Higher-order API where gestures can be tagged with labels and a label can be discerned from a candidate gesture is not provided** - this should be fairly easy to implement by the user of this library and can be tailored to ones needs (e.g. using concurrency, employing some type of quick discard technique).

## Algorithm notes

Following are some remarks regarding the implementation of the gesture matching algorithm, since some aspects differ from the papers.

### Vector normalization

In the `protractor` paper, the vectors are being normalized during scoring. As an optimization, this library normalizes the vectors during pattern construction to reduce calculations during matching.

### Custom distance scoring

One approach for measuring the similarity of two vectors is to sum the distances of the individual points that make up the vector. This is mentioned in the `$P` paper.

Instead, since this library uses a hyperdimensional vector as explained in the `protractor` it also uses the hyperdimensional distance of two vectors to measure similarity.

In order to get a score in the `[0.0..1.0]` range, the following equation is used.

```math
S=1.0 - \frac{distance(template, query)}{2.0}
```

### Custom cosine distance

The `protractor` paper proposes the usage of an inverse cosine distance, as follows:

```math
S = \frac{1.0}{arccos(template \cdot query)}
```

> It is slightly different in the paper, since they do normalization. However, as explained above, this library does normalization beforehand, so the equation is in essence identical.

The problem is that the `arccos` would return values in the range `[0.0...pi]`. And the inverse of that would return values in the range `[0.31..+inf]`. This is inconsistent with other matching functions that return in the range `[0.0..1.0]`.

As such, this library uses a more simplified equation that produces values in the `[0.0..1.0]` range.

```math
S = 1.0 - \frac{arccos(template \cdot query)}{\pi}
```
