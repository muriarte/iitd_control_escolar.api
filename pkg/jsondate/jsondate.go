// jsondate Permite que se puedan procesa fechas en json con el formato 'yyyy-mm-dd'
// ya que normalmente json solo permite fechas con el formato 'yyyy-MM-ddThh:mm:ssZ
// Para lograrlo solo tiene que declarar la variable como tipo JsonDate en vez de
// como time.Time.
// Thanks to https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
package jsondate

import (
	"encoding/json"
	"strings"
	"time"
)

// First create a type alias
type JsonDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonDate) ToTime() time.Time {
	return time.Time(j)
}
