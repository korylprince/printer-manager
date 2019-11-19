package httpapi

import (
	"encoding/json"
	"fmt"
)

func mergeJSON(orig, patch []byte) (interface{}, error) {
	var (
		o interface{}
		p interface{}
	)

	if err := json.Unmarshal(orig, &o); orig != nil && err != nil {
		return nil, fmt.Errorf("Unable unmarshal original json: %v", err)
	}

	if err := json.Unmarshal(patch, &p); patch != nil && err != nil {
		return nil, fmt.Errorf("Unable unmarshal patch json: %v", err)
	}

	return merge(o, p), nil
}

func merge(orig, patch interface{}) interface{} {
	if patch == nil {
		return orig
	}

	if orig == nil {
		return patch
	}

	if origMap, ok := orig.(map[string]interface{}); ok {
		patchMap, ok := patch.(map[string]interface{})
		if !ok {
			return patch
		}

		for k, v := range origMap {
			if v2, ok := patchMap[k]; ok {
				if v2 == nil {
					origMap[k] = nil
				} else {
					origMap[k] = merge(v, v2)
				}
			}
		}

		for k, v := range patchMap {
			if _, ok := origMap[k]; !ok {
				origMap[k] = v
			}
		}

		return origMap
	}

	return patch
}
