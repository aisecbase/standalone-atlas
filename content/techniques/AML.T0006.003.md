---
atlas_id: AML.T0006.003
atlas_type: technique
attack_ref_id: T1595
attack_ref_url: https://attack.mitre.org/techniques/T1595/
created_date: "2026-09-15"
description: Adversaries may send crafted inputs to potential public triggers, such as email addresses, webhooks, or messaging channels, to elicit a response that indicates an invocation of agentic activity. The nature of the...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-09-15"
platforms:
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: Probe AI Agent Trigger Channels
subtechnique_count: 0
subtechnique_of: AML.T0006
tactics:
    - AML.TA0002
title: Probe AI Agent Trigger Channels
url: /techniques/AML.T0006.003/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may send crafted inputs to potential public triggers, such as email addresses, webhooks, or messaging channels, to elicit a response that indicates an invocation of agentic activity. The nature of the response and its content can often reveal agent-managed accounts and expose the additional agentic attack surface. Identified triggers can be used to directly attack the agent.


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
<a class="relation-item" href="/techniques/AML.T0006.001/"><span class="relation-id">AML.T0006.001</span><strong>Запросы к API метаданных платформ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0006.002/"><span class="relation-id">AML.T0006.002</span><strong>Scan for Exposed AI Infrastructure</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0039/"><span class="relation-id">AML.M0039</span><strong>AI Honeypots</strong><p>By capturing adversarial fingerprinting behavior, honeypots can help provide early warning of sweeping activity meant to discover exposed AI targets and attack surfaces being explored by adversaries.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0002 Разведка</span><p>The researchers look for support email addresses on the target organization&#39;s website which may be managed by an AI agent. Then, they probe the system by sending emails and looking for indications of agentic AI in automatic replies.</p></a>
</div>
