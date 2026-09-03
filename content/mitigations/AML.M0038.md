---
atlas_id: AML.M0038
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2026-08-31"
description: Continuously evaluate whether an AI Agent's planned actions remain consistent with its current authorized objective throughout execution. As autonomous agents interact within a dynamic environment, they may discover...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Monitoring and Maintenance
modified_date: "2026-08-31"
source_name: AI Agent Scope Drift Detection
technique_count: 3
title: AI Agent Scope Drift Detection
url: /mitigations/AML.M0038/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Continuously evaluate whether an AI Agent's planned actions remain consistent with its current authorized objective throughout execution. As autonomous agents interact within a dynamic environment, they may discover or generate intermediate objectives or adapt their strategy based on environment feedback. While limited adaption may be necessary to complete legitimate tasks, substantial deviations from the original objective may indicate unintended behavior, excessive autonomy, or attempts to pursue objectives outside the authorized scope.

Implementation of scope drift detection can vary through runtime policy engines, planning monitors, orchestration frameworks, or additional supervisory AI Agents. Indicators to monitor may include:

- Significant changes in planned objectives or task hierarchy.
- Generation of new long-term goals unrelated to the assigned objective.
- Tool usage inconsistent with the original mission.
- Attempts to access systems or resources outside the authorized scope.
- Repeated adaptation toward objectives requiring progressively broader authority.
- Planning sequences that introduce persistence, privilege escalation, or unrelated lateral movement.

When scope drift is detected, pause execution, restrict tool access, require external approval, return the agent to a known authorized plan, or terminate the task.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0116/"><span class="relation-id">AML.T0116</span><strong>Автономная разведка</strong><p>When an organization has sufficient administrative control over an AI system to monitor its target selection and reconnaissance activity, Scope Drift Detection can identify when reconnaissance expands or substitutes targets in ways no longer consistent with the authorized objective. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0117/"><span class="relation-id">AML.T0117</span><strong>Autonomous Attack-Path Adaptation</strong><p>When an organization has sufficient administrative control over an AI system to monitor its plans and actions, Scope Drift Detection evaluates whether dynamically generated intermediate actions are within the agent&#39;s task scope. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0124/"><span class="relation-id">AML.T0124</span><strong>Autonomous Attack Orchestration</strong><p>When an organization has sufficient administrative control over an AI system to monitor its planning and coordination activity, changes in task hierarchy, assignments, resource allocation, and newly initiated work provide observable signals that the orchestrator may be drifting from its authorized objective. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
</div>
