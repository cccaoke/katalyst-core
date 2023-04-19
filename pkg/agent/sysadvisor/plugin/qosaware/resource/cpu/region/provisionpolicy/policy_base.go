/*
Copyright 2022 The Katalyst Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provisionpolicy

import (
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/metacache"
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/plugin/qosaware/resource/cpu/region/regulator"
	"github.com/kubewharf/katalyst-core/pkg/agent/sysadvisor/types"
	"github.com/kubewharf/katalyst-core/pkg/metaserver"
	"github.com/kubewharf/katalyst-core/pkg/metrics"
	"github.com/kubewharf/katalyst-core/pkg/util/machine"
)

type PolicyBase struct {
	regionName string

	requirement int
	regulator   *regulator.CPURegulator

	podSet       types.PodSet
	indicator    types.Indicator
	bindingNumas machine.CPUSet
	essentials   types.ResourceEssentials

	metaReader metacache.MetaReader
	metaServer *metaserver.MetaServer
	emitter    metrics.MetricEmitter
}

func NewPolicyBase(regionName string, regulator *regulator.CPURegulator, metaReader metacache.MetaReader,
	metaServer *metaserver.MetaServer, emitter metrics.MetricEmitter) *PolicyBase {
	cp := &PolicyBase{
		regionName: regionName,
		podSet:     make(types.PodSet),
		indicator:  make(types.Indicator),
		regulator:  regulator,

		metaReader: metaReader,
		metaServer: metaServer,
		emitter:    emitter,
	}
	return cp
}

func (p *PolicyBase) SetPodSet(podSet types.PodSet) {
	p.podSet = podSet.Clone()
}

func (p *PolicyBase) SetIndicator(v types.Indicator) {
	p.indicator = v
}

func (p *PolicyBase) SetRequirement(requirement int) {
	p.requirement = requirement
}

func (p *PolicyBase) SetEssentials(essentials types.ResourceEssentials) {
	p.essentials = essentials
	p.regulator.SetEssentials(essentials)
}

func (p *PolicyBase) SetBindingNumas(numas machine.CPUSet) {
	p.bindingNumas = numas
}