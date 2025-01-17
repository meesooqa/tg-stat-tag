package format

import (
	"github.com/meesooqa/tg-stat-tag/internal/stat"
)

type Formatter interface {
	Format(items []stat.StatItem)
}
