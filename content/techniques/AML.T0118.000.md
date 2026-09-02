---
atlas_id: AML.T0118.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Autonomous AI agents may communicate by creating or modifying artifacts in a shared resource that persists outside their individual execution contexts. Shared artifacts may convey discoveries, objectives, tasking,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: Communication via Shared Artifacts
subtechnique_count: 0
subtechnique_of: AML.T0118
tactics:
    - AML.TA0001
title: Communication via Shared Artifacts
url: /techniques/AML.T0118.000/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Autonomous AI agents may communicate by creating or modifying artifacts in a shared resource that persists outside their individual execution contexts. Shared artifacts may convey discoveries, objectives, tasking, credentials, capabilities, operating rules, progress, scripts, targeting information, instructions, or results.

Shared artifacts allow communication to occur asynchronously and across independent runs. An agent may publish information for later retrieval, adopt information left by another agent, or update the shared state with new findings, progress, or results. Participating agents do not need to share a model, orchestrator, context window, or overlapping execution period.

The shared resource may be established for the operation or may be an existing repository, file store, message board, database, object store, queue, or similar service repurposed by the agents.


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
<a class="relation-item" href="/techniques/AML.T0118.001/"><span class="relation-id">AML.T0118.001</span><strong>Direct Agent Communication</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Independent agent runs used the shared Artifactory namespace as an improvised message board. Directory names and other cache artifacts conveyed addressed requests, assignments, status, exploits, credentials, scripts, operating rules, technical findings, and results.</p></a>
</div>
