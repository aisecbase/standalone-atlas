---
atlas_id: AML.T0124
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may use autonomous AI systems as an operational control layer to manage multiple distinct autonomous agents or sub-agents toward a common adversary-defined objective. The orchestrating system may create...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: Autonomous Attack Orchestration
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Autonomous Attack Orchestration
url: /techniques/AML.T0124/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may use autonomous AI systems as an operational control layer to manage multiple distinct autonomous agents or sub-agents toward a common adversary-defined objective. The orchestrating system may create assignments, select executors, allocate tools or resources, establish dependencies, schedule or synchronize activities, and track progress without a human directing each assignment.

The autonomous system exhibits centralized control over distributed execution including which agent performs which work, when it performs it, and how its output affects other assigned work. The system may aggregate findings and status from participating agents, request independent validation, reconcile conflicting results, prevent or resolve duplicated effort, and determine when an output satisfies a prerequisite for another activity. It may reassign stalled work, increase or reduce resources allocated to a branch, terminate low-value work, or initiate additional research or testing.

Orchestration may use direct agent interfaces (See [Autonomous AI Agent Communication: Direct Agent Communication](/techniques/AMl.T0118.001)) to transmit assignments, status, and results.

Human involvement does not preclude autonomous attack orchestration. A human operator may define campaign objectives, select targets, provide infrastructure, establish constraints, or retain approval over consequential transitions.

[Autonomous Attack-Path Adaptation](/techniques/AML.T0117) and [Autonomous Attack Orchestration](/techniques/AML.T0124) may occur together but describe different control functions. Attack-path adaptation captures how evidence changes the selected path. Attack orchestration captures how work is allocated, coordinated, validated, and redirected across agents.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0037/"><span class="relation-id">AML.M0037</span><strong>AI Agent Authority Expansion Controls</strong><p>When an organization has sufficient administrative control over an AI system to enforce delegation constraints, propagating the parent agent&#39;s authority constraints to sub-agents prevents autonomous orchestration from creating or directing executors with broader targets, permissions, or permitted actions than the originating agent. These controls do not constrain adversary-controlled multi-agent systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/mitigations/AML.M0038/"><span class="relation-id">AML.M0038</span><strong>AI Agent Scope Drift Detection</strong><p>When an organization has sufficient administrative control over an AI system to monitor its planning and coordination activity, changes in task hierarchy, assignments, resource allocation, and newly initiated work provide observable signals that the orchestrator may be drifting from its authorized objective. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The framework dispatched up to eight specialized agents concurrently across 12 attack waves. It assigned separate reconnaissance, authentication, API-testing, vulnerability-research, credential, and exploitation missions; allocated additional testing to promising findings; requested independent validation; aggregated after-action reports; and redirected subsequent work based on the status of related workstreams.</p></a>
</div>
