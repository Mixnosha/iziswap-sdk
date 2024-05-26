package utils

import (
	"fmt"
	"math"
	"math/big"
)

var MIN_TICK *big.Int = big.NewInt(-887272)
var MAX_TICK *big.Int = big.NewInt(887272)


func NearserUsableTick(tick , tickSpacing *big.Int) (*big.Int, error) {
    if tick == nil || tickSpacing == nil {
        return nil, fmt.Errorf("tick or tickSpacing is nil")
    }
    if tickSpacing.Cmp(big.NewInt(0)) <= 0 {
        return nil, fmt.Errorf("tickSpacing is zero or negative")
    }

    if tick.Cmp(MIN_TICK) < 0 {
        return nil, fmt.Errorf("tick is less than MIN_TICK")
    }
    if tick.Cmp(MAX_TICK) > 0 {
        return nil, fmt.Errorf("tick is greater than MAX_TICK")
    }

    ftick := new(big.Float).SetInt(tick)
    ftickSpacing := new(big.Float).SetInt(tickSpacing)

    rounded := new(big.Float).Quo(ftick, ftickSpacing)

    frounded64, _ := rounded.Float64()
    if frounded64 < 0 {
        frounded64 = math.RoundToEven(frounded64)
    }else{
        frounded64 = math.Round(frounded64)
    }

    rounded = new(big.Float).SetFloat64(frounded64)
    rounded.Mul(rounded, ftickSpacing)

    irounded, _ := rounded.Int64()
    brounded := new(big.Int).SetInt64(irounded)

    if brounded.Cmp(MIN_TICK) == -1{
        return brounded.Add(brounded, tickSpacing), nil
    }else if brounded.Cmp(MAX_TICK) == 1{
        return brounded.Sub(brounded, tickSpacing), nil
    }else{
        return brounded, nil
    }
}

