---
atlas_id: AML.TA0015
atlas_type: tactic
attack_ref_id: TA0008
attack_ref_url: https://attack.mitre.org/tactics/TA0008/
created_date: "2025-10-27"
description: Злоумышленник пытается перемещаться по вашей ИИ-среде. Латеральное перемещение включает техники, которые злоумышленники могут использовать для получения доступа к другим системам или компонентам среды и контроля над...
generated: true
generated_by: atlasgen
modified_date: "2025-11-05"
procedure_count: 3
source_name: Lateral Movement
technique_count: 10
title: Латеральное перемещение
url: /tactics/AML.TA0015/
---

Злоумышленник пытается перемещаться по вашей ИИ-среде.

Латеральное перемещение включает техники, которые злоумышленники могут использовать для получения доступа к другим системам или компонентам среды и контроля над ними.
Злоумышленники могут перемещаться в сторону инфраструктуры AI Ops, такой как реестры моделей, системы отслеживания экспериментов, векторные базы данных, ноутбуки или конвейеры обучения.
По мере перемещения по среде злоумышленник может обнаруживать способы доступа к дополнительным инструментам, сервисам или приложениям, связанным с ИИ.
ИИ-агенты также могут быть ценной целью, поскольку обычно имеют больше прав, чем стандартные учетные записи пользователей в системе.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0052/"><span class="relation-id">AML.T0052</span><strong>Фишинг</strong></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091/"><span class="relation-id">AML.T0091</span><strong>Использование альтернативных средств аутентификации</strong></a>
<a class="relation-item" href="/techniques/AML.T0091.000/"><span class="relation-id">AML.T0091.000</span><strong>Токен доступа к приложению</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.000/"><span class="relation-id">AML.T0091.000</span><strong>Токен доступа к приложению</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.001/"><span class="relation-id">AML.T0091.001</span><strong>Cookie веб-сессии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0091.001/"><span class="relation-id">AML.T0091.001</span><strong>Cookie веб-сессии</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0015 Латеральное перемещение</span><p>Злоумышленник использовал извлеченный токен, чтобы аутентифицироваться в backend-сервисе LLM.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее динамические команды, сгенерированные ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0015 Латеральное перемещение</span><p>APT28 отправила с этой учетной записи фишинговое письмо с вложением, содержащим вредоносное ПО.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0015 Латеральное перемещение</span><p>После этого исследователи могли импортировать украденный cookie сеанса сотрудника поддержки в свой браузер, чтобы возобновить аутентифицированный сеанс и потенциально выполнить латеральное перемещение в платформе клиентской поддержки Lenovo от имени этого сотрудника.</p></a>
</div>
