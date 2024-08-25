package attributes

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02 15:04:05"

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", ct.Format(customTimeFormat))
	return []byte(formatted), nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]
	t, err := time.Parse(customTimeFormat, str)
	if err != nil {
		return fmt.Errorf("invalid time format: %v", err)
	}
	ct.Time = t
	return nil
}

// Scan implémente l'interface sql.Scanner, permet de lire une valeur SQL et de la convertir en CustomTime
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime{Time: time.Time{}}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan type %T into CustomTime: %v", value, value)
	}
	*ct = CustomTime{Time: t}
	return nil
}

// Value implémente l'interface driver.Valuer, permet de convertir CustomTime en un type supporté par SQL
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}
