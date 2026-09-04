---
atlas_id: AML.T0102
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-25"
description: Злоумышленники могут использовать большие языковые модели (LLM) для динамической генерации вредоносных команд на основе текста на естественном языке. Обнаруживать динамически сгенерированные команды может быть...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 3
source_name: Generate Malicious Commands
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Генерация вредоносных команд
url: /techniques/AML.T0102/
---

Злоумышленники могут использовать большие языковые модели (LLM) для динамической генерации вредоносных команд на основе текста на естественном языке. Обнаруживать динамически сгенерированные команды может быть сложнее, поскольку сигнатура атаки постоянно меняется. Команды, сгенерированные ИИ, также могут позволить злоумышленникам быстрее адаптироваться к разным средам и корректировать свою тактику.

Злоумышленники могут использовать LLM, находящиеся в среде жертвы, или обращаться к сервисам, размещённым за её пределами. [APT28](https://attack.mitre.org/groups/G0007) использовала модель, размещённую на Hugging Face, в кампании с применением вредоносного ПО LAMEHUG [[logpoint]]. В обоих случаях промпты для генерации вредоносного кода могут сливаться с обычным трафиком.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Блокируйте промпты и выходные данные, которые запрашивают или содержат вредоносные команды.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивайте модели генеративного ИИ в соответствии с целями безопасности, чтобы снизить вероятность генерации ими вредоносных команд или инструкций, способных причинить вред.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Вредоносное ПО LAMEHUG злоупотребляло доступом к модели Qwen 2.5 Coder 32B Instruct через API Hugging Face, чтобы генерировать вредоносные команды из промптов на естественном языке.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The agents generated C code, shell-injection strings, and follow-on shell and Python commands for the exposed harness, revising them as results were returned.</p></a>
<a class="relation-item" href="/studies/AML.CS0070/"><span class="relation-id">AML.CS0070</span><strong>Threat Actor Uses a DeepSeek-Powered Hermes Agent in Langflow and n8n Exploitation Attempts</strong><span class="relation-meta">Актор: Chinese-speaking threat actor using the aliases knaithe and KnYuan / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>DeepSeek generated FOFA queries, shell commands, scanner invocations, and direct HTTP probes based on the results returned during the session.</p></a>
</div>


## Источники

- [LAMEHUG: APT28's First AI-Powered Malware Explained | Guardsix](https://logpoint.com/en/blog/apt28s-new-arsenal-lamehug-the-first-ai-powered-malware)
