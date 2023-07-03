package filters

func GetFilterWrappers() map[string]FilterWrapperCreator {
	return map[string]FilterWrapperCreator{
		"not": NewNotWrapper,
	}
}

func GetFilterBase() map[string]FilterBaseCreator {
	return map[string]FilterBaseCreator{
		// boolean
		"and": NewCompositeAndFilter,
		"or":  NewCompositeOrFilter,
		"not": NewCompositeNotFilter,
		// ip filters
		"ip":             NewIPFilter,
		"geo":            NewGeolocationFilter,
		"reverse_lookup": NewReverseLookupFilter,
		// packet inspection
		"regexp":    NewRegexpFilter,
		"malleable": NewMalleableFilter,
		// misc
		"time": NewTimeFilter,
	}
}