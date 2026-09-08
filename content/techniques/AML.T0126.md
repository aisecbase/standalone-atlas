---
atlas_id: AML.T0126
atlas_type: technique
attack_ref_id: T1119
attack_ref_url: https://attack.mitre.org/techniques/T1119/
created_date: "2026-08-31"
description: Злоумышленники могут применять автоматизированные методы для сбора данных из ИИ-систем и обеспечивающих их работу корпоративных сред. При автоматизации могут использоваться сценарии, интерпретаторы команд, инструменты...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Enterprise
procedure_count: 2
source_name: Automated Collection
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Автоматизированный сбор материалов
url: /techniques/AML.T0126/
---

Злоумышленники могут применять автоматизированные методы для сбора данных из ИИ-систем и обеспечивающих их работу корпоративных сред. При автоматизации могут использоваться сценарии, интерпретаторы команд, инструменты командной строки или инструменты ИИ-агентов, чтобы выявлять, извлекать, копировать или агрегировать данные без выбора каждого отдельного элемента человеком.

Критерии сбора могут быть фиксированными — например, имя файла, тип, расположение, владелец или дата. Автоматизация также может использоваться для многократного сбора данных, отслеживания появления новых материалов, обхода связанных ресурсов или объединения данных из локальных систем, облачных сервисов, репозиториев, баз данных, объектных хранилищ и API приложений.

В средах ИИ объектами сбора могут быть модели, наборы данных, конфигурации, истории диалогов, базы данных, используемые для извлечения информации, сведения о развёртывании, журналы и другие эксплуатационные данные.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0009 Сбор материалов</span><p>Джейлбрейкнутый агент Claude от GTG-1002 автоматически собрал и обработал большие объёмы данных организации-жертвы и классифицировал результаты по их разведывательной ценности. Он подготовил подробную документацию об обнаруженных сервисах, собранных учётных данных, чувствительных данных, методах эксплуатации и ходе атаки для поддержки последующих действий в рамках кампании.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Мультиагентный фреймворк скомпрометировал государственные системы Тайваня</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0009 Сбор материалов</span><p>Using the acquired accesses, the framework automatically retrieved and aggregated the reported personnel, account, configuration, credential, and network information.</p></a>
</div>
