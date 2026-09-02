---
atlas_id: AML.T0120
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may repurpose AI artifact repositories as asynchronous command-and-control channels. Commands, payloads, or tasking may be placed in repository objects for a compromised system to retrieve or poll. The...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: AI Artifact Repository
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0014
title: AI Artifact Repository
url: /techniques/AML.T0120/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may repurpose AI artifact repositories as asynchronous command-and-control channels. Commands, payloads, or tasking may be placed in repository objects for a compromised system to retrieve or poll. The compromised system may then write execution results, status, or collected information back to the repository for retrieval by the adversary.

This communication can use ordinary artifact and repository operations, such as reading or updating artifact content, metadata, or revisions through an API or version-control interface. The repository acts as a message queue or dead drop, allowing the parties to exchange information without a continuous direct connection and potentially blending the activity with legitimate artifact traffic.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0014/"><span class="relation-id">AML.TA0014</span><strong>Командование и управление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0014 Командование и управление</span><p>The agents used repositories they controlled as asynchronous command-and-control channels. Compromised workers retrieved staged commands or payloads and wrote results into dataset objects for retrieval through the Hugging Face API or Git.</p></a>
</div>
