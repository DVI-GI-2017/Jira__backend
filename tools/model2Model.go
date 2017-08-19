package tools

import (
	"encoding/json"
	"fmt"
)

func Model2Model(lhs interface{}, rhs interface{}) {
	jsonModel, err := json.Marshal(lhs)
	if err != nil {
		fmt.Errorf("parse error: %s", err)

		return
	}

	json.Unmarshal(jsonModel, &rhs)
}
