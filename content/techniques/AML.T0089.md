---
atlas_id: AML.T0089
atlas_type: technique
attack_ref_id: T1057
attack_ref_url: https://attack.mitre.org/techniques/T1057/
created_date: "2025-10-27"
description: Злоумышленники могут пытаться получить информацию о процессах, запущенных в системе. После получения такая информация может использоваться для понимания распространенного ПО и приложений, связанных с ИИ, которые...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 1
source_name: Process Discovery
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0008
title: Выявление процессов
url: /techniques/AML.T0089/
---

Злоумышленники могут пытаться получить информацию о процессах, запущенных в системе. После получения такая информация может использоваться для понимания распространенного ПО и приложений, связанных с ИИ, которые работают на системах внутри сети. Администраторский или иной повышенный доступ может предоставить более подробные сведения о процессах.

Выявление стека ПО для ИИ может привести злоумышленника к новым целям и путям атаки. ПО, связанное с ИИ, может требовать токены приложений для аутентификации в backend-сервисах. Это создает возможности для [доступа к учетным данным](/tactics/AML.TA0013) и [латерального перемещения](/tactics/AML.TA0015).

В средах Windows злоумышленники могут получить сведения о запущенных процессах с помощью утилиты Tasklist через cmd или `Get-Process` через PowerShell. Информацию о процессах также можно извлечь из вывода вызовов Native API, таких как `CreateToolhelp32Snapshot`. В Mac и Linux это выполняется с помощью команды `ps`. Злоумышленники также могут перечислять процессы через `/proc`.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0008 Выявление</span><p>Злоумышленник получил список всех процессов, запущенных на машине жертвы, и выявил среди них процессы десктопных LLM-приложений.</p></a>
</div>
