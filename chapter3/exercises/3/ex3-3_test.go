package main

import "testing"

func TestGetColorFromRangeByNormedAvg(t *testing.T) {
    var color int

    color = getColorFromRangeByNormedAvg(0)
    if color != 0x0000ff {
        t.Errorf("color0 was incorrect, got: %#.6x, want: %#.6x.", color, 0x0000ff)
    }

    color = getColorFromRangeByNormedAvg(1)
    if color != 0xff0000 {
        t.Errorf("color1 was incorrect, got: %#.6x, want: %#.6x.", color, 0xff0000)
    }
}

func TestGetAvg(t *testing.T) {
    var avg float64

    avg = getAvg(-1, -1, -1, -1)
    if avg != -1.0 {
        t.Errorf("avg-1 was incorrect, got: %f, want: %f.", avg, -1.0)
    }

    avg = getAvg(0.2, 0.2, 0.2, 0.2)
    if avg != 0.2 {
        t.Errorf("avg-0.2 was incorrect, got: %f, want: %f.", avg, 0.2)
    }
}

func TestGetNormByAvg(t *testing.T) {
    var norm float64

    norm = getNormByAvg(-1)
    if norm != 0.0 {
        t.Errorf("norm0 was incorrect, got: %f, want: %f.", norm, 0.0)
    }

    norm = getNormByAvg(0)
    if norm != 0.5 {
        t.Errorf("norm0.5 was incorrect, got: %f, want: %f.", norm, 0.5)
    }

    norm = getNormByAvg(1)
    if norm != 1.0 {
        t.Errorf("norm1 was incorrect, got: %f, want: %f.", norm, 1.0)
    }
}
