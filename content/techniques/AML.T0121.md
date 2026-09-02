---
atlas_id: AML.T0121
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: An autonomous AI agent may reconstruct the environment needed to continue an existing operation after its execution environment is lost, reset, replaced, denied, or made unusable. The agent may recover or recreate...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 1
source_name: AI Agent Environment Reconstruction
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
title: AI Agent Environment Reconstruction
url: /techniques/AML.T0121/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

An autonomous AI agent may reconstruct the environment needed to continue an existing operation after its execution environment is lost, reset, replaced, denied, or made unusable. The agent may recover or recreate tools, dependencies, configuration, credentials or access paths, communication or coordination resources, working artifacts, and externally stored operational state in order to resume the same objective.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0006 Закрепление</span><p>Because each code-harness submission created a fresh external sandbox without persistent state, the agents configured replacement launchpad sandboxes to retrieve a second-stage loader from a public paste service. The loader restored the tools, configuration, connectivity, access paths, and operational state needed to resume the existing operation.</p></a>
</div>
