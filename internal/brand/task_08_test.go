package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask08(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	b := BrothBatch{ID: "b", BoneLotIDs: []string{"lot"}, WaterLiters: 40, BoneKilograms: 6}
	lots := map[string]IngredientLot{"lot": {ID: "lot", IngredientCode: "yak-bone"}}
	require.NoError(t, s.CheckBroth(context.Background(), b, lots))
}
