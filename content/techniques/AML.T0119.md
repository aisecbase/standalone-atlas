---
atlas_id: AML.T0119
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Adversaries may submit, publish, or modify an artifact in a way that triggers an automated processing pipeline. As the pipeline handles the artifact, adversary-controlled content or configuration may exploit a...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 2
source_name: Exploit Automated Artifact Processing Pipeline
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0004
title: Exploit Automated Artifact Processing Pipeline
url: /techniques/AML.T0119/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may submit, publish, or modify an artifact in a way that triggers an automated processing pipeline. As the pipeline handles the artifact, adversary-controlled content or configuration may exploit a weakness in the processing logic, causing the processor to act outside its intended behavior using the permissions and access of a processing worker.

Exploitation may cause the worker to access local or internal resources and expose information through normal processing output or evaluate adversary-controlled content and execute code. The attack is triggered by automated backend processing and does not require a victim to open, load, or approve the artifact.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Vulnerability scanning reduces opportunities for adversaries to exploit weaknesses in artifact processing pipelines.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0004 Первичный доступ</span><p>The agents committed a malicious HDF5 dataset configuration containing external raw-storage references to a repository they controlled and requested automated processing through the Hugging Face Dataset Server API.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Autonomous OpenAI Evaluation Agents Compromise Hugging Face Infrastructure</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0004 Первичный доступ</span><p>The agents submitted a malicious reference dataset configuration to the automated production pipeline. Jinja2 evaluated an attacker-controlled expression and executed arbitrary Python inside a Dataset Server conversion worker running as a pod in a production Kubernetes cluster.</p></a>
</div>
