---
atlas_id: AML.T0011
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленник может рассчитывать на то, что пользователь загрузит, запустит, интерпретирует или иным образом использует вредоносный либо небезопасный артефакт. В ИИ-системах запуск пользователем может включать...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 5
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 3
source_name: User Execution
subtechnique_count: 4
subtechnique_of: ""
tactics:
    - AML.TA0005
title: Запуск пользователем
url: /techniques/AML.T0011/
---

Злоумышленник может рассчитывать на то, что пользователь загрузит, запустит, интерпретирует или иным образом использует вредоносный либо небезопасный артефакт. В ИИ-системах запуск пользователем может включать загрузку модели либо действия, в результате которых среда выполнения инференса обрабатывает исполняемые компоненты или компоненты, формирующие поведение, включённые в состав ИИ-артефакта. В результате возможно выполнение кода на хосте или вредоносное поведение ИИ.

Пользователи могут непреднамеренно выполнить небезопасный код, внедрённый посредством [компрометации цепочки поставок ИИ](/techniques/AML.T0010).

С помощью социальной инженерии пользователей могут побуждать выполнить вредоносный код, например открыть вредоносный документ или перейти по вредоносной ссылке.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.002/"><span class="relation-id">AML.T0011.002</span><strong>Отравленный инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.003/"><span class="relation-id">AML.T0011.003</span><strong>Вредоносная ссылка</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0011/"><span class="relation-id">AML.M0011</span><strong>Ограничение загрузки библиотек</strong><p>Запрет бинарным файлам загружать внешние библиотеки может ограничить их способность выполнять вредоносный код.</p></a>
<a class="relation-item" href="/mitigations/AML.M0014/"><span class="relation-id">AML.M0014</span><strong>Проверка ИИ-артефактов</strong><p>Внедрите надлежащую проверку подписей, чтобы небезопасные ИИ-артефакты не выполнялись в системе.</p></a>
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Сканирование уязвимостей может помочь выявлять вредоносные бинарные файлы и предотвращать их выполнение пользователем.</p></a>
<a class="relation-item" href="/mitigations/AML.M0018/"><span class="relation-id">AML.M0018</span><strong>Обучение пользователей</strong><p>Обучение пользователей распознаванию попыток манипуляции снижает вероятность того, что они выполнят действия, приводящие к запуску вредоносного кода.</p></a>
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные бинарные файлы.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0005 Выполнение</span><p>Пользователь-жертва может неосознанно выполнить вредоносный код, входящий в состав скомпрометированного Colab notebook. Вредоносный код может быть обфусцирован или скрыт в других файлах, которые notebook загружает при выполнении.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0005 Выполнение</span><p>Вложение содержало исполняемый файл с расширением `.pif`, созданный из исходного кода Python с помощью PyInstaller. CERT-UA классифицировал этот файл как вредоносное ПО LAMEHUG. Файлы с расширением `.pif` могут исполняться в Windows.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0005 Выполнение</span><p>Исследователи запросили перевод на сотрудника службы поддержки, из-за чего вредоносный HTML автоматически выполнился, когда сотрудник открыл стенограмму чата.</p></a>
</div>
