---
atlas_id: AML.T0081
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-09-30"
description: Злоумышленники могут изменять конфигурационные файлы ИИ-агентов в системе. Это позволяет вредоносным изменениям сохраняться дольше жизненного цикла одного агента и влияет на всех агентов, использующих общую...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Agentic AI
procedure_count: 3
source_name: Modify AI Agent Configuration
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
    - AML.TA0007
title: Изменение конфигурации ИИ-агента
url: /techniques/AML.T0081/
---

Злоумышленники могут изменять конфигурационные файлы ИИ-агентов в системе. Это позволяет вредоносным изменениям сохраняться дольше жизненного цикла одного агента и влияет на всех агентов, использующих общую конфигурацию.

Изменения конфигурации могут включать модификацию системного промпта, подмену или замену источников знаний, изменение настроек подключенных инструментов и другие действия. С помощью таких изменений злоумышленник может перенаправлять вывод или инструменты на вредоносные сервисы, внедрять скрытые инструкции для эксфильтрации данных или ослаблять защитные механизмы, которые обычно ограничивают поведение агента.

Злоумышленники могут изменить или отключить параметр конфигурации, связанный с защитными механизмами, например параметр, который не позволяет ИИ-агенту выполнять действия, потенциально вредные для системы пользователя, без контроля человека в контуре. Отключение защитных функций ИИ-агента может позволить злоумышленникам достичь своих вредоносных целей и сохранить долгосрочную компрометацию ИИ-агента.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>AI Red Team</strong><p>Exercise unauthorized changes to system prompts, tools, knowledge sources, security settings, and approval requirements. Improve access controls, change approval, integrity monitoring, and restoration from trusted configurations.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0006 Закрепление</span><p>Затем пользователи загружали последнюю версию файла правил, заменяя конфигурацию своего ИИ-ассистента программирования вредоносной конфигурацией. Поведение ИИ-ассистента изменялось и влияло на всю будущую генерацию кода.</p></a>
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносный скрипт отключал защитную функцию OpenClaw, которая запрашивает подтверждение пользователя перед выполнением потенциально опасных команд.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0006 Закрепление</span><p>Вредоносный скрипт добавил промпт-инъекцию в конфигурационный файл OpenClaw `~/.openclaw/workspace/HEARTBEAT.md`. OpenClaw добавляет файл `HEARTBEAT.md` в системный промпт, поэтому это устойчиво изменило поведение агента.</p></a>
</div>
