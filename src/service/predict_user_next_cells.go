package service

import (
	"math"
	"sort"

	"github.com/kajiLabTeam/mr-platform-recommend-contents-server/common"
	"github.com/uber/h3-go/v4"
)

func PredictUserNextCells(lat, lon float64, level int) []h3.Cell {
	latLng := h3.NewLatLng(lat, lon)

	// 最も近い 3件の Resolution が level のIDを取得
	cellResolution10 := h3.LatLngToCell(latLng, level)
	h3GridDisks := h3.GridDisk(cellResolution10, 1)

	// 周辺のセルを取得して、距離を測るための構造体
	var cellDistances []common.CellDistance
	for _, v := range h3GridDisks {
		// 2点間の距離を求める
		lat1 := latLng.Lat
		lon1 := latLng.Lng
		tmpLatLng := v.LatLng()
		lat2 := tmpLatLng.Lat
		lon2 := tmpLatLng.Lng
		// 緯度経度の差を求める
		latDiff := lat1 - lat2
		lonDiff := lon1 - lon2
		// 2点間の距離を求める
		distance := math.Pow(latDiff, 2) + math.Pow(lonDiff, 2)

		cellDistances = append(cellDistances, common.CellDistance{Cell: v, Distance: distance})
	}

	// ソートして3番目より大きいものは削除
	sort.Slice(cellDistances, func(i, j int) bool { return cellDistances[i].Distance < cellDistances[j].Distance })
	cellDistances = cellDistances[:2]

	var cells []h3.Cell
	for _, v := range cellDistances {
		cells = append(cells, v.Cell)
	}

	return cells
}
