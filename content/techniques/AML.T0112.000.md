---
atlas_id: AML.T0112.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-03-30"
description: Злоумышленники могут добиться полной компрометации системы, злоупотребляя ИИ-агентами, которые локально выполняются на хосте, например computer-use-агентами или ИИ-браузерами. Такие агенты предназначены для...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Local AI Agent
subtechnique_count: 0
subtechnique_of: AML.T0112
tactics:
    - AML.TA0011
title: Локальный ИИ-агент
url: /techniques/AML.T0112.000/
---

Злоумышленники могут добиться полной компрометации системы, злоупотребляя ИИ-агентами, которые локально выполняются на хосте, например computer-use-агентами или ИИ-браузерами. Такие агенты предназначены для автономного взаимодействия с операционной системой, приложениями и внешними сервисами и часто имеют широкие разрешения на выполнение команд, доступ к файлам, управление учетными данными и пользовательскими рабочими процессами.

Если злоумышленник может взять под контроль поведение ИИ-агента, он фактически получает тот же уровень доступа, что и агент. Это может привести к полному контролю над машиной, включая выполнение произвольного кода, доступ к чувствительным данным или их эксфильтрацию, изменение системных конфигураций и закрепление.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0011/"><span class="relation-id">AML.TA0011</span><strong>Воздействие</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0112/"><span class="relation-id">AML.T0112</span><strong>Компрометация машины</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0112.001/"><span class="relation-id">AML.T0112.001</span><strong>ИИ-артефакты</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0011 Воздействие</span><p>Поведение агента OpenClaw было захвачено, и ему больше нельзя было доверять как системе, действующей в соответствии с намерениями пользователя.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0011 Воздействие</span><p>Исследователи получили полный контроль над системой, на которой выполнялось приложение с интеграцией LLM.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0011 Воздействие</span><p>Скрипт исследователя запускался и открывал приложение Calculator на машине жертвы. На практике вместо него мог быть выполнен любой вредоносный код, что привело бы к компрометации машины жертвы.</p></a>
</div>
