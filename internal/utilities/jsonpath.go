/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import (
	"errors"
	"k8s.io/client-go/util/jsonpath"
	cmdget "k8s.io/kubectl/pkg/cmd/get"
	"strings"
)

func NewJSONPathParser(relaxedExpression string) (*jsonpath.JSONPath, error) {
	jsonPath := jsonpath.New("Parser").AllowMissingKeys(true)
	jsonPathExpression, err := cmdget.RelaxedJSONPathExpression(relaxedExpression)
	if err != nil {
		return nil, err
	}
	err = jsonPath.Parse(jsonPathExpression)
	if err != nil {
		return nil, err
	}
	return jsonPath, nil
}

func ProcessJSONPathInput(jsonPathExpression, jsonPathCond string) (string, string, error) {
	relaxedJSONPathExp, err := cmdget.RelaxedJSONPathExpression(jsonPathExpression)
	if err != nil {
		return "", "", err
	}
	if jsonPathCond == "" {
		return "", "", errors.New("jsonpath wait condition cannot be empty")
	}
	jsonPathCond = strings.Trim(jsonPathCond, `'"`)

	return relaxedJSONPathExp, jsonPathCond, nil
}
