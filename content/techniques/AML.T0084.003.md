---
atlas_id: AML.T0084.003
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут извлекать цепочки вызовов из конфигураций ИИ-агентов, что может раскрыть потенциальные цели для удаленного выполнения кода (RCE) или других уязвимостей. Уязвимые цепочки вызовов часто связывают...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 1
source_name: Call Chains
subtechnique_count: 0
subtechnique_of: AML.T0084
tactics:
    - AML.TA0008
title: Цепочки вызовов
url: /techniques/AML.T0084.003/
---

Злоумышленники могут извлекать цепочки вызовов из конфигураций ИИ-агентов, что может раскрыть потенциальные цели для удаленного выполнения кода (RCE) или других уязвимостей. Уязвимые цепочки вызовов часто связывают пользовательский ввод или вывод LLM с точкой выполнения, например `exec`, `eval` или `os.popen`. Позднее эти уязвимости могут быть использованы через [промпт-инъекцию LLM](/techniques/AML.T0051).

Злоумышленники могут систематически выявлять потенциально уязвимые цепочки вызовов, присутствующие во фреймворках LLM, а затем сканировать приложения, настроенные на использование этих цепочек вызовов, чтобы выбрать цели [[arxiv]].


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0084/"><span class="relation-id">AML.T0084</span><strong>Выявление конфигурации ИИ-агента</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0084.000/"><span class="relation-id">AML.T0084.000</span><strong>Встроенные знания</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Limit public disclosure of agent call chains and execution-sink details.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0008 Выявление</span><p>Исследователи использовали статический анализ, чтобы извлечь цепочки вызовов из исходного кода целевых приложений и определить, какие из них используют функции LLM-фреймворков, уязвимые к RCE.</p></a>
</div>


## Источники

- [\[2309.02926\] Demystifying RCE Vulnerabilities in LLM-Integrated Apps](https://arxiv.org/abs/2309.02926)
