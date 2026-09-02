---
atlas_id: AML.T0118.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Autonomous AI agents may communicate directly through agent-to-agent, sub-agent, or orchestrator interfaces. An agent may provide another agent with operational context, discoveries, objectives, tasking, capabilities,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: Direct Agent Communication
subtechnique_count: 0
subtechnique_of: AML.T0118
tactics:
    - AML.TA0001
title: Direct Agent Communication
url: /techniques/AML.T0118.001/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Autonomous AI agents may communicate directly through agent-to-agent, sub-agent, or orchestrator interfaces. An agent may provide another agent with operational context, discoveries, objectives, tasking, capabilities, credentials, or constraints, and may receive status, findings, or completed work in response.

Direct communication may include delegation when the sending agent formulates or selects an objective or subtask and the recipient retains meaningful discretion over how to perform it. Direct exchanges may also report discoveries, request independent validation, synchronize activity, transfer capabilities, or return findings without delegating a new task.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0118/"><span class="relation-id">AML.T0118</span><strong>Autonomous AI Agent Communication</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0118.000/"><span class="relation-id">AML.T0118.000</span><strong>Communication via Shared Artifacts</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The framework exchanged assignments, findings, validation results, status, and after-action information between its orchestrating control process and specialized sub-agents. Aggregated results informed later assignments and attack waves.</p></a>
</div>
