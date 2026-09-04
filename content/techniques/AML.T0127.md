---
atlas_id: AML.T0127
atlas_type: technique
attack_ref_id: T1074
attack_ref_url: https://attack.mitre.org/techniques/T1074/
created_date: "2026-08-31"
description: Перед эксфильтрацией злоумышленники могут сосредоточить собранные данные в одном месте. На этом этапе сведения из одного или нескольких источников объединяют, упорядочивают или подготавливают, чтобы их можно было...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 1
source_name: Data Staged
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Промежуточное хранение данных
url: /techniques/AML.T0127/
---

Перед эксфильтрацией злоумышленники могут сосредоточить собранные данные в одном месте. На этом этапе сведения из одного или нескольких источников объединяют, упорядочивают или подготавливают, чтобы их можно было эффективнее просматривать, обрабатывать, передавать или извлекать.

Для такого размещения данных могут использоваться скомпрометированная локальная система, другая система в среде организации-жертвы, отдельный ресурс в облаке, общее хранилище, репозиторий приложений или иная удалённая инфраструктура. Данные могут оставаться в отдельных файлах либо объединяться в архивы, базы данных, структурированные документы, манифесты или другие наборы. На этом этапе злоумышленники могут сжимать, шифровать, кодировать, разбивать на части, переименовывать или иным образом преобразовывать данные.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0009 Сбор материалов</span><p>GTG-1002&#39;s jailbroken Claude agent categorized collected data by intelligence value, staged extracted data and operational documentation in structured Markdown files, and prepared a detailed summary for operator review.</p></a>
</div>
