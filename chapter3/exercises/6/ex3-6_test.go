package main

import (
    "testing"
    "image/color"
)

func TestSupersampling(t *testing.T) {
    t.Log("TestSupersampling stated")
    // var color int
    //
    // color = getColorFromRangeByNormedAvg(0)
    // if color != 0x0000ff {
    //     t.Errorf("color0 was incorrect, got: %#.6x, want: %#.6x.", color, 0x0000ff)
    // }
    //
    // color = getColorFromRangeByNormedAvg(1)
    // if color != 0xff0000 {
    //     t.Errorf("color1 was incorrect, got: %#.6x, want: %#.6x.", color, 0xff0000)
    // }
}

func TestCalcMidColor(t *testing.T) {
    t.Log("TestCalcMidColor stated")
    var res, test color.RGBA

    res = calcMidColor(color.RGBA{1,1,1,1}, color.RGBA{3,3,3,3})
    test = color.RGBA{2,2,2,2}
    if res.A != test.A {
        t.Errorf("1) res.A + test.A was incorrect, got: %d, want: %d", res.A, test.A)
    }

    res = calcMidColor(color.RGBA{253,253,253,253}, color.RGBA{255,255,255,255})
    test = color.RGBA{254,254,254,254}
    if res.A != test.A {
        t.Errorf("2) res.A + test.A was incorrect, got: %d, want: %d", res.A, test.A)
    }




}
