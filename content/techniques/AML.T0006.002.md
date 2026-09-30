---
atlas_id: AML.T0006.002
atlas_type: technique
attack_ref_id: T1595
attack_ref_url: https://attack.mitre.org/techniques/T1595/
created_date: "2026-09-15"
description: Adversaries may scan network ports and services to identify deployed AI backends, model-serving endpoints, and AI agent infrastructure reachable over the internet. Self-hosted AI runtimes typically listen on...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 0
source_name: Scan for Exposed AI Infrastructure
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Scan for Exposed AI Infrastructure
url: /techniques/AML.T0006.002/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may scan network ports and services to identify deployed AI backends, model-serving endpoints, and AI agent infrastructure reachable over the internet. Self-hosted AI runtimes typically listen on predictable, well-known ports and expose standard API paths, allowing adversaries to efficiently locate candidate hosts and infer the platform running on them.

After identifying candidate hosts, adversaries can interact with them directly to confirm live AI services, fingerprint the software stack, and determine version and configuration details. This information can be used to select exploitable targets and tailor subsequent access attempts.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006/"><span class="relation-id">AML.T0006</span><strong>Активное сканирование</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0006.000/"><span class="relation-id">AML.T0006.000</span><strong>Enumerate Hosted AI Resources</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Query Platform Metadata APIs</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Probe AI Agent Trigger Channels</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
</div>
