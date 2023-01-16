package shared

func CompareAndPatchIfEmptyString(checkedValue string, patchValue string) string {
	if checkedValue == "" {
		return patchValue
	}
	return checkedValue
}
