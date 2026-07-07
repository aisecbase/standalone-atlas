---
atlas_id: AML.T0011.003
atlas_type: technique
attack_ref_id: T1204
attack_ref_url: https://attack.mitre.org/techniques/T1204/
created_date: "2026-01-30"
description: Злоумышленник может рассчитывать на то, что пользователь перейдет по вредоносной ссылке, чтобы добиться выполнения кода. Пользователя могут подвергнуть социальной инженерии и заставить перейти по ссылке, которая...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Malicious Link
subtechnique_count: 0
subtechnique_of: AML.T0011
tactics:
    - AML.TA0005
title: Вредоносная ссылка
url: /techniques/AML.T0011.003/
---

Злоумышленник может рассчитывать на то, что пользователь перейдет по вредоносной ссылке, чтобы добиться выполнения кода. Пользователя могут подвергнуть социальной инженерии и заставить перейти по ссылке, которая приведет к выполнению кода. Такое действие пользователя обычно наблюдается как последующее поведение после перехода по ссылке из целевого фишингового сообщения. Переход по ссылке также может приводить к другим техникам выполнения, например эксплуатации уязвимости браузера или приложения на стороне клиента. Ссылки также могут приводить пользователя к скачиванию вредоносных файлов, требующих запуска.

Злоумышленник может разными способами использовать вредоносные ссылки, чтобы получить доступ к системе организации-жертвы через ИИ-систему. Например, ИИ-агент, настроенный без проверки заголовков источника веб-сайта, будет принимать подключения с любого сайта, что позволяет злоумышленникам обходить ранее недоступные сетевые ограничения.


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
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0011.002/"><span class="relation-id">AML.T0011.002</span><strong>Отравленный инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0005 Выполнение</span><p>Когда жертва переходила по ссылке на сайт исследователя, вредоносный JavaScript-скрипт выполнялся в ее браузере.</p></a>
</div>
