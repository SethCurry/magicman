package mtg

import (
	"fmt"
	"strings"
)

var (
	ColorRed = Color{
		Full:  "Red",
		Short: "R",
	}
	ColorBlue = Color{
		Full:  "Blue",
		Short: "U",
	}
	ColorBlack = Color{
		Full:  "Black",
		Short: "B",
	}
	ColorWhite = Color{
		Full:  "White",
		Short: "W",
	}
	ColorGreen = Color{
		Full:  "Green",
		Short: "G",
	}

	Colors = []Color{
		ColorRed,
		ColorBlue,
		ColorBlack,
		ColorWhite,
		ColorGreen,
	}
)

type Color struct {
	Full  string
	Short string
}

func (c *Color) UnmarshalJSON(contents []byte) error {
	switch strings.Trim(string(contents), "\"") {
	case ColorRed.Full, ColorRed.Short:
		*c = ColorRed
		return nil
	case ColorBlue.Full, ColorBlue.Short:
		*c = ColorBlue
		return nil
	case ColorBlack.Full, ColorBlack.Short:
		*c = ColorBlack
		return nil
	case ColorWhite.Full, ColorWhite.Short:
		*c = ColorWhite
		return nil
	case ColorGreen.Full, ColorGreen.Short:
		*c = ColorGreen
		return nil
	}

	return fmt.Errorf("unknown color code: %s", string(contents))
}
