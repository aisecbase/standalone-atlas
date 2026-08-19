---
atlas_id: AML.T0102
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-25"
description: Злоумышленники могут использовать большие языковые модели (LLM), чтобы динамически генерировать вредоносные команды из запросов на естественном языке. Динамически сгенерированные команды может быть сложнее обнаружить,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Generate Malicious Commands
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Генерация вредоносных команд
url: /techniques/AML.T0102/
---

Злоумышленники могут использовать большие языковые модели (LLM), чтобы динамически генерировать вредоносные команды из запросов на естественном языке. Динамически сгенерированные команды может быть сложнее обнаружить, поскольку сигнатура атаки постоянно меняется. Команды, сгенерированные ИИ, также могут позволить злоумышленникам быстрее адаптироваться к разным средам и корректировать свою тактику.

Злоумышленники могут использовать LLM, присутствующие в среде жертвы, или обращаться к сервисам, размещенным вне среды жертвы. [APT28](https://attack.mitre.org/groups/G0007) использовала модель, размещенную на Hugging Face, в кампании с вредоносным ПО LAMEHUG [[logpoint]]. В обоих случаях промпты для генерации вредоносного кода могут смешиваться с обычным трафиком.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Block prompts and outputs that request or contain malicious commands.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Align generative AI models with safety objectives to reduce the likelihood that they will generate malicious commands or harmful instructions.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее динамические команды, сгенерированные ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>LAMEHUG злоупотреблял Hugging Face API модели Qwen 2.5 Coder 32B Instruct, чтобы генерировать вредоносные команды из промптов на естественном языке.</p></a>
</div>


## Источники

- [LAMEHUG: APT28's First AI-Powered Malware Explained | Guardsix](https://logpoint.com/en/blog/apt28s-new-arsenal-lamehug-the-first-ai-powered-malware)
