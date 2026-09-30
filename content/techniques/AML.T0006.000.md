---
atlas_id: AML.T0006.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may directly probe an agentic or SaaS platform to enumerate the resources a specific victim has deployed on it. Platforms frequently host AI agents behind predictable URL structures derived from...
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
source_name: Enumerate Hosted AI Resources
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Enumerate Hosted AI Resources
url: /techniques/AML.T0006.000/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may directly probe an agentic or SaaS platform to enumerate the resources a specific victim has deployed on it. Platforms frequently host AI agents behind predictable URL structures derived from identifiers such as environment, tenant, resource-group, or agent names, and default out-of-the-box deployment configurations keep these conventions consistent across victims. Adversaries can learn these conventions from public sources such as vendor documentation, code repositories and actual hosted resources, then fuzz or brute-force the derived namespace to discover live AI agents, endpoints, and associated metadata. Discovered resources can be used to identify targets for further access, collection, or attack adaptation.


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
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Query Platform Metadata APIs</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Scan for Exposed AI Infrastructure</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Probe AI Agent Trigger Channels</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
</div>
