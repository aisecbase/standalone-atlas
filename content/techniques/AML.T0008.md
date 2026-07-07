---
atlas_id: AML.T0008
atlas_type: technique
attack_ref_id: T1583
attack_ref_url: https://attack.mitre.org/techniques/T1583/
created_date: "2021-05-13"
description: Злоумышленники могут покупать, арендовать или брать во временное пользование инфраструктуру для использования на протяжении своей операции. Для размещения и оркестрации операций злоумышленников существует широкий...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 4
source_name: Acquire Infrastructure
subtechnique_count: 6
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Получение инфраструктуры
url: /techniques/AML.T0008/
---

Злоумышленники могут покупать, арендовать или брать во временное пользование инфраструктуру для использования на протяжении своей операции.

Для размещения и оркестрации операций злоумышленников существует широкий спектр инфраструктуры.

Инфраструктурные решения включают физические или облачные серверы, домены, мобильные устройства и сторонние веб-сервисы.

Бесплатные ресурсы также могут использоваться, но обычно они ограничены.

Инфраструктура также может включать физические компоненты, например средства противодействия, которые ухудшают работу или нарушают функционирование компонентов ИИ или датчиков, включая печатные материалы, носимые устройства или маскировку.

Использование таких инфраструктурных решений позволяет злоумышленнику подготавливать, запускать и выполнять операцию.

Такие решения могут помогать операциям злоумышленника сливаться с трафиком, который воспринимается как нормальный, например с обращениями к сторонним веб-сервисам.

В зависимости от реализации злоумышленники могут использовать инфраструктуру, которую трудно физически связать с ними, а также инфраструктуру, которую можно быстро развернуть, изменить и отключить.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0008.000/"><span class="relation-id">AML.T0008.000</span><strong>Рабочие пространства для разработки ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.001/"><span class="relation-id">AML.T0008.001</span><strong>Потребительское оборудование</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.002/"><span class="relation-id">AML.T0008.002</span><strong>Домены</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.003/"><span class="relation-id">AML.T0008.003</span><strong>Физические средства противодействия</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.004/"><span class="relation-id">AML.T0008.004</span><strong>Serverless-инфраструктура</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.005/"><span class="relation-id">AML.T0008.005</span><strong>Прокси для ИИ-сервисов</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0029/"><span class="relation-id">AML.CS0029</span><strong>Эксфильтрация разговоров Google Bard</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь установил, что Google Apps Script можно вызвать через URL на `script.google.com` или `googleusercontent.com` и настроить так, чтобы аутентификация не требовалась. Это позволяет вызвать скрипт без срабатывания Content Security Policy Bard.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи приобрели домен `aisystem.tech` для размещения вредоносного скрипта и промптов.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход защитных ограничений Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Storm-2139 приобрела инфраструктуру для поддержки сервиса, который продавал доступ к Azure OpenAI Service с обходом защитных ограничений.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи развернули сервер для получения чувствительной информации, эксфильтрированной из уязвимого LLM-сервиса.</p></a>
</div>
