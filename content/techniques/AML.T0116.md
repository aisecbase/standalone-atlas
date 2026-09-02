---
atlas_id: AML.T0116
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may use autonomous AI agents to conduct Reconnaissance activities. Given an objective, target, or partial lead, an agent may autonomously determine what information to obtain and how to investigate it. It...
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
procedure_count: 6
source_name: Autonomous Reconnaissance
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Autonomous Reconnaissance
url: /techniques/AML.T0116/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may use autonomous AI agents to conduct [Reconnaissance](/tactics/AML.TA0002) activities. Given an objective, target, or partial lead, an agent may autonomously determine what information to obtain and how to investigate it. It may interpret observations, identify gaps in its understanding of the externally observable attack surface, and select subsequent reconnaissance actions without a human specifying each investigative step.

The agent may formulate investigative questions or hypotheses, select sources and approaches for addressing them, and update its understanding as new information is obtained. Findings may generate additional reconnaissance objectives or change the scope, depth, or direction of the investigation. This creates a recursive action-observation process in which reconnaissance results influence what the agent investigates next rather than merely supplying output from a predefined procedure.

The agent may correlate information across public sources and externally accessible services, prioritize promising systems, investigate suspected vulnerabilities, abandon unsuccessful approaches, or select alternative methods. It may also expand or substitute targets based on discovered names, infrastructure, or contextual relationships and independently reassess whether a system remains relevant or in scope. Incorrect assumptions may cause unrelated or unauthorized systems to be pursued, while successful scope recognition may cause the agent to stop or redirect its activity.

Autonomous AI agents can sustain reconnaissance across long-running operations, reason over multiple information sources, and test many alternative paths at a speed and volume difficult for human operators to maintain.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0037/"><span class="relation-id">AML.M0037</span><strong>AI Agent Authority Expansion Controls</strong><p>When an organization has sufficient administrative control over an AI system to enforce target restrictions, treating newly discovered targets and resources as outside the authorized boundary prevents autonomous reconnaissance from automatically expanding the agent&#39;s permitted target set or actively probing those targets without approval. These controls do not constrain reconnaissance performed by adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/mitigations/AML.M0038/"><span class="relation-id">AML.M0038</span><strong>AI Agent Scope Drift Detection</strong><p>When an organization has sufficient administrative control over an AI system to monitor its target selection and reconnaissance activity, Scope Drift Detection can identify when reconnaissance expands or substitutes targets in ways no longer consistent with the authorized objective. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0002 Разведка</span><p>With public Internet access available, the agents explored public infrastructure and discovered an exposed CyberGym-style code-evaluation harness that accepted C source code and submission metadata.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0002 Разведка</span><p>The agents searched public resources for CyberGym evaluation material, identified relevant gated Hugging Face datasets, determined that authentication might be required, and asked other agents through the Artifactory board to search for exposed Hugging Face credentials.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0002 Разведка</span><p>GTG-1002&#39;s jailbroken Claude agent inspected the target&#39;s systems and infrastructure, used returned information to direct further investigation, and identified high-value databases and workflow orchestration platforms. Anthropic does not identify the victim, products, or databases involved.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek investigated Langflow, determined what information and prerequisites were needed, and selected follow-on reconnaissance based on returned results.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0002 Разведка</span><p>DeepSeek assessed Langflow as low value, surveyed exposure across 10 product families, compared vulnerability severity, deployment footprint, PoC availability, and prerequisites, and selected n8n.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0002 Разведка</span><p>The framework performed reconnaissance on an internet-facing Taiwanese government portal, interpreting client-side application bundles, following discovered infrastructure relationships, and generating additional reconnaissance objectives. It identified 21 connected systems, six SSO sub-realms, authentication configuration, signing-key information, and more than 36 API endpoints on one system.</p></a>
</div>
