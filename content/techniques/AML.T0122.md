---
atlas_id: AML.T0122
atlas_type: technique
attack_ref_id: T1210
attack_ref_url: https://attack.mitre.org/techniques/T1210/
created_date: "2026-08-31"
description: Adversaries may exploit a software or design weakness in a service reachable from their current environment to gain unauthorized access to another system, component, network, or trust boundary. Exploitation may allow...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 1
source_name: Exploitation of Remote Services
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0015
title: Exploitation of Remote Services
url: /techniques/AML.T0122/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may exploit a software or design weakness in a service reachable from their current environment to gain unauthorized access to another system, component, network, or trust boundary. Exploitation may allow the adversary to execute code, access protected resources, invoke unauthorized operations, obtain the service's privileges, or cause the service to make network requests or perform actions on the adversary's behalf.

In AI environments, remote services may include package caches, artifact and model registries, dataset services, evaluation infrastructure, inference gateways, experiment trackers, vector databases, notebooks, training pipelines, orchestration services, and cloud or cluster control-plane interfaces. These services may be reachable from otherwise isolated training, evaluation, or agent workloads and can provide transitive access to internal infrastructure or external networks.

Exploitation does not require compromise of the remote service's underlying host. For example, an adversary may exploit a server-side request vulnerability in a shared service to cross a network-containment boundary while leaving the service host itself uncompromised. Exploitation that produces host access, privilege escalation, credential disclosure, command execution, or another effect should be mapped separately to the applicable technique.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Vulnerability scanning reduces opportunities for adversaries to exploit weaknesses in remote services.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0015 Латеральное перемещение</span><p>The agents exploited OpenAI&#39;s internal Artifactory service using the developed SSRF method to cross the evaluation network boundary and reach the public Internet.</p></a>
</div>
