---
atlas_id: AML.T0011.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут разрабатывать небезопасные ИИ-артефакты, которые при запуске оказывают вредоносное воздействие. Злоумышленник может использовать эту технику для закрепления постоянного доступа к системам. Такие...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 6
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 4
source_name: Unsafe AI Artifacts
subtechnique_count: 0
subtechnique_of: AML.T0011
tactics:
    - AML.TA0005
title: Небезопасные ИИ-артефакты
url: /techniques/AML.T0011.000/
---

Злоумышленники могут разрабатывать небезопасные ИИ-артефакты, которые при запуске оказывают вредоносное воздействие.

Злоумышленник может использовать эту технику для закрепления постоянного доступа к системам.

Такие модели могут внедряться через [компрометацию цепочки поставок ИИ](/techniques/AML.T0010).

Сериализация моделей является распространенным методом хранения, передачи и загрузки моделей.

Однако без надлежащей проверки этот формат создает возможность для выполнения кода.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011/"><span class="relation-id">AML.T0011</span><strong>Запуск пользователем</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.002/"><span class="relation-id">AML.T0011.002</span><strong>Отравленный инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.003/"><span class="relation-id">AML.T0011.003</span><strong>Вредоносная ссылка</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0011/"><span class="relation-id">AML.M0011</span><strong>Ограничение загрузки библиотек</strong><p>Ограничьте загрузку библиотек ML-артефактами.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Предотвращайте выполнение ML-артефактов, которые не подписаны должным образом.</p></a>
<a class="relation-item" href="/mitigations/AML.M0014/"><span class="relation-id">AML.M0014</span><strong>Проверка ИИ-артефактов</strong><p>Внедрите надлежащую проверку подписей, чтобы небезопасные ИИ-артефакты не выполнялись в системе.</p></a>
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные ИИ-артефакты, такие как модели или данные, и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/mitigations/AML.M0018/"><span class="relation-id">AML.M0018</span><strong>Обучение пользователей</strong><p>Обучайте пользователей распознавать попытки манипуляции, чтобы они не запускали небезопасный код, который при выполнении может создать небезопасные артефакты. Такие артефакты могут негативно повлиять на систему.</p></a>
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные артефакты моделей.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0005 Выполнение</span><p>Когда любой пользователь позже загрузит модель, она автоматически выполнит полезную нагрузку злоумышленника.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0005 Выполнение</span><p>Если пользователь загружал вредоносную модель, выполнялась вредоносная нагрузка злоумышленника.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0005 Выполнение</span><p>Жертва загружает и использует отравленный артефакт в совместимом движке инференса. Во время инференса движок автоматически интерпретирует включённый в состав артефакта шаблон чата, из-за чего изменённая злоумышленником логика формирования промпта выполняется в рамках обычного использования модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Model Namespace Reuse Supply Chain Attack</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0005 Выполнение</span><p>When a user or service deployed the malicious model, loading or deployment executed the embedded payload in the model endpoint environment.</p></a>
</div>
