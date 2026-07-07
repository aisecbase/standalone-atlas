---
atlas_id: AML.T0051.002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-11-04"
description: Злоумышленник может запускать промпт-инъекцию через действие пользователя или событие, происходящее в среде жертвы. Триггерные промпт-инъекции часто нацелены на ИИ-агентов, которые могут активироваться способами,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Triggered
subtechnique_count: 0
subtechnique_of: AML.T0051
tactics:
    - AML.TA0005
title: Триггерная промпт-инъекция
url: /techniques/AML.T0051.002/
---

Злоумышленник может запускать промпт-инъекцию через действие пользователя или событие, происходящее в среде жертвы. Триггерные промпт-инъекции часто нацелены на ИИ-агентов, которые могут активироваться способами, выявленными злоумышленником на этапе [выявления](/tactics/AML.TA0008) (см. [триггеры активации](/techniques/AML.T0084.002)).

Такие вредоносные промпты могут быть скрыты или обфусцированы от пользователя и уже находиться где-то в среде жертвы после выполнения злоумышленником техники [инфильтрации промпта через публично доступное приложение](/techniques/AML.T0093). Этот тип инъекции может использоваться злоумышленником для закрепления в системе или для атаки на неосведомленного пользователя системы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051.000/"><span class="relation-id">AML.T0051.000</span><strong>Прямая промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/mitigations/AML.M0033/"><span class="relation-id">AML.M0033</span><strong>Валидация входных и выходных данных компонентов ИИ-агента</strong><p>Валидация может помешать злоумышленникам выполнять промпт-инъекции, способные повлиять на агентные рабочие процессы.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0005 Выполнение</span><p>Когда почтовый ассистент извлекает письмо с червем при очередной генерации ответа, промпт-инъекция меняет поведение GenAI-почтового ассистента.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0005 Выполнение</span><p>Исследователи получают ответ на указанный ими адрес, что указывает на наличие ИИ-агента и успешное срабатывание триггерной промпт-инъекции.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0005 Выполнение</span><p>Когда позднее пользователь вызвал Copilot, тот извлек вредоносное письмо в свой контекст, что привело к срабатыванию промпт-инъекции.</p></a>
</div>
