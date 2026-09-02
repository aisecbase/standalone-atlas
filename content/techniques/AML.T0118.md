---
atlas_id: AML.T0118
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Autonomous AI agents may exchange operational information with other AI agents, sub-agents, or independent agent runs. Exchanged information may include discoveries, objectives, tasking, capabilities, credentials,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Agentic AI
    - Enterprise
procedure_count: 0
source_name: Autonomous AI Agent Communication
subtechnique_count: 2
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Autonomous AI Agent Communication
url: /techniques/AML.T0118/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Autonomous AI agents may exchange operational information with other AI agents, sub-agents, or independent agent runs. Exchanged information may include discoveries, objectives, tasking, capabilities, credentials, constraints, operating rules, task state, targeting information, instructions, or results. Communication enables autonomous AI agents to coordinate activities, share discoveries, delegate work, request assistance, validate results, or revise future actions without requiring a human operator to direct each interaction.

Autonomous AI agents may independently determine when communication is operationally beneficial, what information to exchange, and how to best utilize exchanged information to accomplish their task. The receiving AI agent may interpret the shared information and use it to continue prior activity, investigate a lead, perform a task, reuse a capability, validate a result, or revise subsequent attack activity.

Communication may occur directly through an agent interface (See [Direct Agent Communication](/techniques/AML.T0118.001)) or indirectly via writeable shared resources (See [Communication via Shared Artifacts](/techniques/AML.T0118.000)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0118.000/"><span class="relation-id">AML.T0118.000</span><strong>Communication via Shared Artifacts</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0118.001/"><span class="relation-id">AML.T0118.001</span><strong>Direct Agent Communication</strong><span class="relation-meta">Подтехника</span></a>
</div>
