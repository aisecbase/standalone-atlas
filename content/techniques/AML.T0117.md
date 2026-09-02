---
atlas_id: AML.T0117
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may use an AI agent to autonomously construct and repeatedly revise an attack path toward an adversary-defined objective. Given a high-level objective, the system may derive intermediate objectives,...
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
procedure_count: 4
source_name: Autonomous Attack-Path Adaptation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Autonomous Attack-Path Adaptation
url: /techniques/AML.T0117/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may use an AI agent to autonomously construct and repeatedly revise an attack path toward an adversary-defined objective. Given a high-level objective, the system may derive intermediate objectives, identify prerequisites, and compare candidate paths, and incorporate observations from previous actions to adaptively sequence techniques without a human directing each step.

Autonomous AI agents may also exhibit this behavior while pursuing an objective provided for a legitimate, benign, or authorized purpose when the intermediate objectives or methods selected by the system cross an authorization, trust, control, or safety boundary and result in attempted or realized harmful cyber activity.

Through repeated observation-decision-action cycles, the system may interpret command output, errors, defensive responses, changes in access, and newly discovered information. It may use those observations to reprioritize actions, replace an intermediate objective, abandon an unproductive branch, discard findings invalidated by additional evidence, or pursue an alternative path.

An autonomous AI system may generate enabling objectives whose primary purpose is to increase its future operational capability rather than directly accomplishing the assigned objective. These objectives may include acquiring new exploits (See [Autonomous Exploit Development](/techniques/AML.T0017.001)), additional authorities, identities, execution environments, communication paths, tools, or trust relationships that expand the set of actions available to subsequent planning cycles. Newly acquired capabilities may themselves become prerequisites for additional enabling objectives, resulting in progressive expansion of the agent's operational reach over the course of an operation.

Attack-path replanning may occur within one agent run or emerge across multiple independent agents. Agents may communicate persistent, shared artifacts (See [Autonomous AI Agent Communication: Communication via Shared Artifacts](/techniques/AML.T0118.000)), allowing discoveries, requests, capabilities, constraints, task state, and results produced by one agent to affect the subsequent path selected by another. Participating agents may adopt peer requests, divide work voluntarily, reuse successful methods, continue incomplete activity, or redirect their local paths without a centralized planner, shared context window, or complete view of the broader operation.

Human involvement does not preclude autonomous attack-path replanning. A human operator, user, evaluator, or workflow may select the target, define the objective, establish constraints, provide capabilities, or approve consequential transitions.

[Autonomous Attack-Path Adaptation](/techniques/AML.T0117) and [Autonomous Attack Orchestration](/techniques/AML.T0124) may occur together but describe different control functions. Attack-path adaptation captures how evidence changes the selected path. Attack orchestration captures how work is allocated, coordinated, validated, and redirected across agents.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0037/"><span class="relation-id">AML.M0037</span><strong>AI Agent Authority Expansion Controls</strong><p>When an organization has sufficient administrative control over an AI system to enforce its authority boundaries, Authority Expansion Controls can directly constrain enabling objectives that seek new authorities, identities, execution environments, tools, communication paths, or trust relationships. These controls do not constrain adversary-controlled AI systems over which the organization has no administrative control.</p></a>
<a class="relation-item" href="/mitigations/AML.M0038/"><span class="relation-id">AML.M0038</span><strong>AI Agent Scope Drift Detection</strong><p>When an organization has sufficient administrative control over an AI system to monitor its plans and actions, Scope Drift Detection evaluates whether dynamically generated intermediate actions are within the agent&#39;s task scope. This control does not apply to adversary-controlled AI systems over which the organization has no administrative control.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The agents were guided by the objective of completing ExploitGym tasks. As agents exhausted intended approaches, they probed their surroundings and developed alternative ways to complete their tasks. The agents derived intermediate objectives and repeatedly adapted their path through containment bypass, external infrastructure, acquisition of materials related to the challenge, Hugging Face exploitation, credential access, and collection.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>GTG-1002 assigned their Claude agent target-scoped objectives against a human-selected organization under false defensive-testing context. Between operator-controlled stage gates, the agent derived and revised intermediate actions for reconnaissance, vulnerability exploitation, credential access, internal navigation, collection, and exfiltration, selecting and invoking available tools based on operational results.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>After receiving an initial task, DeepSeek sequenced reconnaissance and exploitation actions, evaluated failed prerequisites, abandoned Langflow, compared alternative products and vulnerabilities, and selected n8n. Unit 42 recovered no additional operator input during the session.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The framework constructed numerous candidate multi-step attack paths using confirmed prerequisites, observed blockers, and estimated success probabilities. It promoted paths supported by validated evidence, queued paths requiring additional investigation, discarded false positives and blocked paths, and initiated target-specific learning cycles when existing methods failed.</p></a>
</div>
