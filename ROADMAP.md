# Roadmap

## HeightField

Change it to use squares, so then we compute the acceleration from the square
the trace is in and neighboring traces. We could use all 8 neighbors to improve
accuracy, weighting the diagonals accordingly since they're further. This is
simpler than using points. Also, we wouldn't have to worry as much about the
edges. We could also weight the acceleration by the spot within the square if
that became desirable.

## Field

Add a helper to use a height map directly, computing vectors based on it. This
will make computed landscapes simpler, but also support creating fields from
topographical data.

## Renderer

The `Renderer` should provide better flexibility for drawing the sim history in
various ways. The user should be able to provide a callback that receives the
full `TraceFrame` and returns a color. This should, perhaps, be an interface so
that line style and other drawing properties can also be determined. A set of
default implementations should exist. For example, one that chooses a color
based on velocity. To support these complex color use cases, we should also
provide a `ColorScale` type of some sort that can map values on a range into
colors.

We should change `Renderer` to be an interface that allows the user to specify
a history along with the various drawing parameters, such as those described
above. The Ebiten interface implementation can be a private implementation
detail of that renderer.

Some of the things described above could be provided now by a custom `Drawer`,
but that is, perhaps, too general and complicated for some common use cases,
like color, so we should determine which use cases need simple options and
provide those, then provide `Drawer` implementation as a fallback for more
advanced use cases.
