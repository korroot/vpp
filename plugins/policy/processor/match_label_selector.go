/*
 * // Copyright (c) 2017 Cisco and/or its affiliates.
 * //
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at:
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

package processor

import (
	nsmodel "github.com/contiv/vpp/plugins/ksr/model/namespace"
	podmodel "github.com/contiv/vpp/plugins/ksr/model/pod"
	policymodel "github.com/contiv/vpp/plugins/ksr/model/policy"
	"github.com/contiv/vpp/plugins/policy/utils"
)

const (
	in           = policymodel.Policy_LabelSelector_LabelExpression_IN
	notIn        = policymodel.Policy_LabelSelector_LabelExpression_NOT_IN
	exists       = policymodel.Policy_LabelSelector_LabelExpression_EXISTS
	doesNotExist = policymodel.Policy_LabelSelector_LabelExpression_DOES_NOT_EXIST
)

// isPodLabelSelectorMatch finds if podLabelSelecor matches Pod's labels or not
func (pp *PolicyProcessor) isPodLabelSelectorMatch(
	pod *podmodel.Pod,
	matchPodLabels []*policymodel.Policy_Label,
	matchPodExpressions []*policymodel.Policy_LabelSelector_LabelExpression,
	policyNamespace string) bool {

	if len(matchPodLabels) > 0 && len(matchPodExpressions) > 0 {
		if !pp.isPodLabelMatch(pod, matchPodLabels, policyNamespace) ||
			!pp.isPodExpressionMatch(pod, matchPodExpressions, policyNamespace) {
			return false
		}
		return true

	} else if len(matchPodLabels) == 0 && len(matchPodExpressions) > 0 {
		if !pp.isPodExpressionMatch(pod, matchPodExpressions, policyNamespace) {
			return false
		}
		return true

	} else if len(matchPodLabels) > 0 && len(matchPodExpressions) == 0 {
		if !pp.isPodLabelMatch(pod, matchPodLabels, policyNamespace) {
			return false
		}
		return true

	}

	return true
}

// isPodLabelMatch finds if pod labels match a collection of pod selector labels.
func (pp *PolicyProcessor) isPodLabelMatch(pod *podmodel.Pod,
	matchPodLabels []*policymodel.Policy_Label, policyNamespace string) bool {

	podNamespace := pod.Namespace
	podLabels := pod.Label
	labelExists := make(map[string]bool)
	labelPrefix := podNamespace + "/"

	for _, podLabel := range podLabels {
		keyLabel := labelPrefix + podLabel.Key + "/" + podLabel.Value
		labelExists[keyLabel] = true
	}

	return isMatchLabel(matchPodLabels, labelExists, labelPrefix)
}

// isPodExpressionMatch finds if pod labels match a collection of pod selector expressions.
func (pp *PolicyProcessor) isPodExpressionMatch(pod *podmodel.Pod,
	matchPodExpressions []*policymodel.Policy_LabelSelector_LabelExpression, policyNamespace string) bool {

	podNamespace := pod.Namespace
	podLabels := pod.Label
	podMap := make(map[string]string)

	for _, podLabel := range podLabels {
		podMapKey := podNamespace + "/" + podLabel.Key
		podMap[podMapKey] = podLabel.Value
	}
	expressionKeyPrefix := podNamespace + "/"

	return isMatchExpression(matchPodExpressions, podMap, expressionKeyPrefix)
}

// isNsLabelSelectorMatch finds if namespace LabelSelector matches Pod's namespace labels.
func (pp *PolicyProcessor) isNsLabelSelectorMatch(
	pod *podmodel.Pod,
	matchNsLabels []*policymodel.Policy_Label,
	matchNsExpressions []*policymodel.Policy_LabelSelector_LabelExpression) bool {

	if len(matchNsLabels) > 0 && len(matchNsExpressions) > 0 {
		if !pp.isNsLabelMatch(pod, matchNsLabels) ||
			!pp.isNsExpressionMatch(pod, matchNsExpressions) {
			return false
		}
		return true

	} else if len(matchNsLabels) == 0 && len(matchNsExpressions) > 0 {
		if !pp.isNsExpressionMatch(pod, matchNsExpressions) {
			return false
		}
		return true

	} else if len(matchNsLabels) > 0 && len(matchNsExpressions) == 0 {
		if !pp.isNsLabelMatch(pod, matchNsLabels) {
			return false
		}
		return true

	}

	return true
}

// isNsLabelMatch finds if pod's namespace labels match a collection of namespace selector labels.
func (pp *PolicyProcessor) isNsLabelMatch(pod *podmodel.Pod,
	matchNsLabels []*policymodel.Policy_Label) bool {
	//Get Pod's namespace labels
	podNamespaceID := nsmodel.ID(pod.Namespace)
	_, nsData := pp.Cache.LookupNamespace(podNamespaceID)
	nsLabels := nsData.Label
	labelExists := make(map[string]bool)

	for _, nsLabel := range nsLabels {
		label := nsLabel.Key + "/" + nsLabel.Value
		labelExists[label] = true
	}
	labelPrefix := ""

	return isMatchLabel(matchNsLabels, labelExists, labelPrefix)
}

// isNsExpressionMatch finds if pod's namespace labels match a collection of namespace selector expressions.
func (pp *PolicyProcessor) isNsExpressionMatch(pod *podmodel.Pod,
	matchNsExpressions []*policymodel.Policy_LabelSelector_LabelExpression) bool {
	//Get Pod's namespace labels
	podNamespaceID := nsmodel.ID(pod.Namespace)
	_, nsData := pp.Cache.LookupNamespace(podNamespaceID)
	nsLabels := nsData.Label
	nsMap := make(map[string]string)

	for _, nsLabel := range nsLabels {
		nsMap[nsLabel.Key] = nsLabel.Value
	}
	expressionKeyPrefix := ""
	return isMatchExpression(matchNsExpressions, nsMap, expressionKeyPrefix)

}

// isNsUpdateLabelSelectorMatch finds if pod's namespace labels match a collection of namespace selector labels,
// after an update in namespace labels
func (pp *PolicyProcessor) isNsUpdateLabelSelectorMatch(
	ns *nsmodel.Namespace,
	matchNsUpdateLabels []*policymodel.Policy_Label,
	matchNsUpdateExpressions []*policymodel.Policy_LabelSelector_LabelExpression) bool {

	if len(matchNsUpdateLabels) > 0 && len(matchNsUpdateExpressions) > 0 {
		if !pp.isNsUpdateLabelMatch(ns, matchNsUpdateLabels) ||
			!pp.isNsUpdateExpressionMatch(ns, matchNsUpdateExpressions) {
			return false
		}
		return true

	} else if len(matchNsUpdateLabels) == 0 && len(matchNsUpdateExpressions) > 0 {
		if !pp.isNsUpdateExpressionMatch(ns, matchNsUpdateExpressions) {
			return false
		}
		return true

	} else if len(matchNsUpdateLabels) > 0 && len(matchNsUpdateExpressions) == 0 {
		if !pp.isNsUpdateLabelMatch(ns, matchNsUpdateLabels) {
			return false
		}
		return true

	}

	return true
}

// isNsUpdateLabelMatch finds if namespace LabelSelector matches Pod's namespace labels,
// after an update in namespace labels
func (pp *PolicyProcessor) isNsUpdateLabelMatch(ns *nsmodel.Namespace,
	matchNsUpdateLabels []*policymodel.Policy_Label) bool {
	//Get Pod's namespace labels
	_, nsData := pp.Cache.LookupNamespace(nsmodel.GetID(ns))
	nsLabels := nsData.Label
	labelExists := make(map[string]bool)

	for _, nsLabel := range nsLabels {
		label := nsLabel.Key + "/" + nsLabel.Value
		labelExists[label] = true
	}
	labelPrefix := ""
	return isMatchLabel(matchNsUpdateLabels, labelExists, labelPrefix)

}

// isNsUpdateExpressionMatch finds if pod's namespace labels match a collection of namespace selector expressions,
// after an update in namespace labels
func (pp *PolicyProcessor) isNsUpdateExpressionMatch(ns *nsmodel.Namespace,
	matchNsUpdateExpressions []*policymodel.Policy_LabelSelector_LabelExpression) bool {
	//Get Pod's namespace labels
	_, nsData := pp.Cache.LookupNamespace(nsmodel.GetID(ns))
	nsLabels := nsData.Label
	nsMap := make(map[string]string)

	for _, nsLabel := range nsLabels {
		nsMap[nsLabel.Key] = nsLabel.Value
	}
	expressionKeyPrefix := ""
	return isMatchExpression(matchNsUpdateExpressions, nsMap, expressionKeyPrefix)

}

func isMatchLabel(matchLabels []*policymodel.Policy_Label, labelExists map[string]bool, matchLabelPrefix string) bool {

	isMatch := true
	for _, matchLabel := range matchLabels {
		label := matchLabelPrefix + matchLabel.Key + "/" + matchLabel.Value
		if labelExists[label] != true {
			isMatch = false
			break
		}
	}

	return isMatch
}

func isMatchExpression(matchExpressions []*policymodel.Policy_LabelSelector_LabelExpression,
	keyMap map[string]string, expressionKeyPrefix string) bool {

	isMatch := false
	for _, matchExpression := range matchExpressions {
		switch matchExpression.Operator {
		case in:
			labels := utils.ConstructLabels(matchExpression.Key, matchExpression.Value)
			for _, label := range labels {
				if keyMap[label.Key] == label.Value {
					isMatch = true
					break
				} else {
					isMatch = false
				}
			}
			if !isMatch {
				return false
			}

		case notIn:
			labels := utils.ConstructLabels(matchExpression.Key, matchExpression.Value)
			for _, label := range labels {
				if keyMap[label.Key] == label.Value {
					return false
				} else {
					isMatch = true
				}
				isMatch = true
			}

		case exists:
			expressionKey := expressionKeyPrefix + matchExpression.Key
			if keyMap[expressionKey] != "" {
				isMatch = true
			} else {
				return false
			}

		case doesNotExist:
			expressionKey := expressionKeyPrefix + matchExpression.Key
			if keyMap[expressionKey] == "" {
				return false
			} else {
				isMatch = true
			}
			isMatch = true
		}
	}

	return isMatch
}
