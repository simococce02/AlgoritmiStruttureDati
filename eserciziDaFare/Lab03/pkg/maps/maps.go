package maps

import "constraints"

// SumValues will return the total of all the maps values
func SumValues[T comparable, U constraints.Integer](source map[T]U) U {
	var total U
	for _, val := range source {
		total += val
	}
	return total
}

// SumValuesFor will return the total of all the maps values where the keys match the predicate
func SumValuesFor[T comparable, U constraints.Integer](source map[T]U, predicate func(T) bool) U {
	var total U
	for key, val := range source {
		if predicate(key) {
			total += val
		}
	}
	return total
}

// MaxKey will return the largest key
func MaxKey[T constraints.Ordered, U any](source map[T]U) T {
	started := false
	var maxKey T
	for key := range source {
		if !started {
			maxKey = key
			started = true
		}
		if key > maxKey {
			maxKey = key
		}
	}
	return maxKey
}

// MaxValue will return the entry with the largest value
func MaxValue[T comparable, U constraints.Integer](source map[T]U) (T, U) {
	started := false
	var maxKey T
	var maxVal U
	for key, val := range source {
		if !started {
			maxKey = key
			maxVal = val
			started = true
		}
		if val > maxVal {
			maxKey = key
			maxVal = val
		}
	}
	return maxKey, maxVal
}

// MinValue will return the entry with the smallest value
func MinValue[T comparable, U constraints.Integer](source map[T]U) (T, U) {
	return MinMappedValue(source, func(u U) U { return u })
}

// MinMappedValue will return the entry with the smallest value returned from func
func MinMappedValue[T comparable, U any, V constraints.Ordered](source map[T]U, op func(U) V) (T, U) {
	started := false
	var minKey T
	var minVal U
	var minCompare V
	for key, val := range source {
		if !started {
			minKey = key
			minVal = val
			minCompare = op(val)
			started = true
		}
		testCompare := op(val)
		if testCompare < minCompare {
			minKey = key
			minVal = val
			minCompare = testCompare
		}
	}
	return minKey, minVal
}

// ContainsKey will return true if the provided key is in the map
func ContainsKey[T comparable, U any](source map[T]U, key T) bool {
	for k := range source {
		if k == key {
			return true
		}
	}
	return false
}

// First will find the first key value pair that matches the provided predicate
func First[T comparable, U any](source map[T]U, predicate func(k T, v U) bool) (T, U, bool) {
	for k, v := range source {
		if predicate(k, v) {
			return k, v, true
		}
	}
	var blankKey T
	var blankVal U
	return blankKey, blankVal, false
}

// AnyKey returns an unspecified key from the map
func AnyKey[T comparable, U any](source map[T]U) T {
	for k := range source {
		return k
	}
	var blankKey T
	return blankKey
}

// Keys returns all keys in the map
func Keys[T comparable, U any](source map[T]U) []T {
	var result []T
	for k := range source {
		result = append(result, k)
	}
	return result
}
