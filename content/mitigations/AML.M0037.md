---
atlas_id: AML.M0037
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2026-08-31"
description: Limit an AI agent's ability to autonomously acquire, assume, or otherwise obtain additional authorities that expand its effective permissions during execution. The maximum authority available to the agent should be...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
    - Monitoring and Maintenance
modified_date: "2026-08-31"
source_name: AI Agent Authority Expansion Controls
technique_count: 3
title: AI Agent Authority Expansion Controls
url: /mitigations/AML.M0037/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Limit an AI agent's ability to autonomously acquire, assume, or otherwise obtain additional authorities that expand its effective permissions during execution. The maximum authority available to the agent should be explicitly granted prior to runtime. Additional resources, identities, services, and targets discovered during execution should be treated as outside the authorized boundary unless they are independently validated and added to scope. All authority expansion controls should be implemented outside the AI agent and should not rely solely on system prompts, model alignment, or the agent recognizing that an action is out of scope. Implementations of these controls may be achieved through enforcement mechanisms such as:

- Policy engines
- Target allowlists
- Protocol and destination restrictions
- Approval gates
- Preventing the agent from using credentials that were not approved for the task
- Monitoring and auditing changes in the agent's effective authority over time

Authority expansion controls include placing restrictions on the number, scope, duration, and concurrent use of authentication and/or authorization tokens available during execution. Tokens may include API access tokens, OAuth tokens, cloud IAM session credentials, service account tokens, Git tokens, or other short-lived authentication artifacts. When policy limits are reached or exceeded, organizations may revoke access, prevent additional token acquisition, require human approval, or terminate the agent's execution.

Propagate the original authority constraints to sub-agents and delegated tasks. A delegated agent may receive narrower restrictions but should not expand the parent agent's scope, authority, targets, or permitted actions.

Authority expansion controls should be implemented alongside permissions configurations for AI agents and tools (See [Privileged AI Agent Permissions Configuration](/mitigations/AML.M0026), [Single-User AI Agent Permissions Configuration](/mitigations/AML.M0027), [AI Agent Tools Permissions Configuration](/mitigations/AML.M0028)). Attempted changes in scope should be accompanied with [Human In-the-Loop for AI Agent Actions](/mitigations/AML.M0029). Log new resource discovery, denials, exceptions, approvals, and scope changes using [AI Telemetry Logging](/mitigations/AML.M0024).


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0116/"><span class="relation-id">AML.T0116</span><strong>Autonomous Reconnaissance</strong><p>When an organization has sufficient administrative control over an AI system to enforce target restrictions, treating newly discovered targets and resources as outside the authorized boundary prevents autonomous reconnaissance from automatically expanding the agent&#39;s permitted target set or actively probing those targets without approval. These controls do not constrain reconnaissance performed by adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0117/"><span class="relation-id">AML.T0117</span><strong>Autonomous Attack-Path Adaptation</strong><p>When an organization has sufficient administrative control over an AI system to enforce its authority boundaries, Authority Expansion Controls can directly constrain enabling objectives that seek new authorities, identities, execution environments, tools, communication paths, or trust relationships. These controls do not constrain adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/techniques/AML.T0124/"><span class="relation-id">AML.T0124</span><strong>Autonomous Attack Orchestration</strong><p>When an organization has sufficient administrative control over an AI system to enforce delegation constraints, propagating the parent agent&#39;s authority constraints to sub-agents prevents autonomous orchestration from creating or directing executors with broader targets, permissions, or permitted actions than the originating agent. These controls do not constrain adversary-controlled multi-agent systems over which the organization has no administrative control.</p></a>
</div>
