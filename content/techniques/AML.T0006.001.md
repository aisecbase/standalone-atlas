---
atlas_id: AML.T0006.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-09-15"
description: Adversaries may query documented or undocumented APIs in agentic SaaS hosting platforms to uncover agentic targets. SaaS and agentic platforms can expose provider control-plane or identity APIs that return tenant,...
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
source_name: Query Platform Metadata APIs
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Query Platform Metadata APIs
url: /techniques/AML.T0006.001/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may query documented or undocumented APIs in agentic SaaS hosting platforms to uncover agentic targets. SaaS and agentic platforms can expose provider control-plane or identity APIs that return tenant, environment, or deployment information IDs, including for resources that are misconfigured or unintentionally accessible by unauthenticated users.

Attackers have been seen abusing this type of functionality to perform information gathering on SaaS platforms, for example via AADInternals' OSINT page,[[aadinternals]] an OSINT online tool showcasing an undocumented API reconnaissance method for Entra ID. This undocumented Power Platform API could be used to uncover environment IDs, which may subsequently be used to scan for public agents. Following the identified abuse, required authentication was added to the tool.


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
<a class="relation-item" href="/techniques/AML.T0006.000/"><span class="relation-id">AML.T0006.000</span><strong>Получение списка размещённых ресурсов ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Scan for Exposed AI Infrastructure</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.003/"><span class="relation-id">AML.T0006.003</span><strong>Probe AI Agent Trigger Channels</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
</div>
