---
atlas_id: AML.T0115
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Adversaries may create or modify AI artifacts and publish them through public or shared distribution channels to facilitate compromise of downstream AI systems. Poisoned AI artifacts may include datasets, models, and...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 3
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Publish Poisoned AI Artifacts
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Publish Poisoned AI Artifacts
url: /techniques/AML.T0115/
---

> Перевод описания пока не добавлен; ниже показан оригинальный текст ATLAS.

Adversaries may create or modify AI artifacts and publish them through public or shared distribution channels to facilitate compromise of downstream AI systems. Poisoned AI artifacts may include datasets, models, and AI agent tools containing malicious content, behaviors, code, or configurations.

Adversaries may publish novel artifacts or malicious variants of legitimate artifacts through dataset or model repositories, package registries, source code repositories, tool hubs, or remotely hosted services. Victims may subsequently acquire and integrate these artifacts through [AI Supply Chain Compromise](/techniques/AML.T0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115.000/"><span class="relation-id">AML.T0115.000</span><strong>Datasets</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.001/"><span class="relation-id">AML.T0115.001</span><strong>Models</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.002/"><span class="relation-id">AML.T0115.002</span><strong>AI Agent Tools</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Dataset repositories inspect submissions and quarantine poisoned samples, labels, annotations, or metadata before listing.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Model repositories evaluate submissions for backdoors, data leakage, adversarial influence, and unexpected behavior before listing.</p></a>
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Model and agent tool registries scan uploaded artifacts for malicious content before listing.</p></a>
</div>
