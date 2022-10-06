package art

import (
	"github.com/glesica/drifter/internal/history"
	"image/color"
)

type Drawer func(iter *history.Iterator)

func LineDrawer(canvas Canvas) Drawer {
	return func(iter *history.Iterator) {
		canvas.SetColor(color.Black)

		ok, snapshot := iter.Next()
		if !ok {
			return
		}

		canvas.MoveTo(snapshot.X(), snapshot.Y())

		for {
			ok, snapshot := iter.Next()
			if !ok {
				break
			}

			canvas.LineTo(snapshot.X(), snapshot.Y())
		}

		canvas.Stroke()
	}
}