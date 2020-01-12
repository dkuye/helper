package helper

import "net/url"

func ValidationFieldError(vErrors url.Values, fieldName string) string {
	_, ok := vErrors[fieldName]
	if !ok {
		return ""
	}
	return vErrors[fieldName][0]
}

func ValidationString(s []string) string {
	//strings.Replace(s, "[", "", -1)
	//strings.Replace(s, "]", "", -1)
	return s[0]
}

func FormValues(formValues map[string][]string, fieldName string) string {
	_, ok := formValues[fieldName]
	if !ok {
		return ""
	}
	return formValues[fieldName][0]

}
