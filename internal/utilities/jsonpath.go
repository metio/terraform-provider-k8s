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

func newJSONPathParser(jsonPathExpression string) (*jsonpath.JSONPath, error) {
	j := jsonpath.New("wait")
	if jsonPathExpression == "" {
		return nil, errors.New("jsonpath expression cannot be empty")
	}
	if err := j.Parse(jsonPathExpression); err != nil {
		return nil, err
	}
	return j, nil
}

func processJSONPathInput(jsonPathExpression, jsonPathCond string) (string, string, error) {
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
