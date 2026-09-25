package parser

import (
	"encoding/binary"
	"time"

	"github.com/avaksru/ble2homed/pkg/types"
)

// parseMiBandFEE0Data разбирает Service Data с UUID 0xFEE0
// (Huami/Amazfit fitness service, используется Mi Band 4 и др.).
//
// Формат (эмпирически подтверждён на реальных пакетах Mi Band 4):
//
//	data[0:2] — количество шагов, uint16 little-endian
//	data[2:]  — padding (нули)
//
// Примеры реальных пакетов (Service Data для UUID FEE0):
//
//	0D 04 00 00  →  0x040D = 1037 шагов
//	30 04        →  0x0430 = 1072 шагов
func parseMiBandFEE0Data(data []byte, now time.Time) map[string]types.ParsedValue {
	result := make(map[string]types.ParsedValue)

	if len(data) < 2 {
		return result
	}

	steps := binary.LittleEndian.Uint16(data[0:2])

	result["steps"] = types.ParsedValue{
		Value:     int(steps),
		Unit:      "steps",
		Type:      "steps",
		Source:    "Mi Band FEE0",
		Timestamp: now,
	}

	return result
}