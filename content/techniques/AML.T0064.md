---
atlas_id: AML.T0064
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут выявлять источники данных, используемые в системах RAG (генерации, дополненной извлечением), для выбора целей атаки. Точно определив эти источники, они могут сосредоточиться на отравлении или иной...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Gather RAG-Indexed Targets
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Сбор целей, индексируемых RAG
url: /techniques/AML.T0064/
---

Злоумышленники могут выявлять источники данных, используемые в системах RAG (генерации, дополненной извлечением), для выбора целей атаки. Точно определив эти источники, они могут сосредоточиться на отравлении или иной манипуляции внешними репозиториями данных, на которые опирается ИИ.

Данные, индексируемые RAG, могут быть выявлены в публичной документации системы или через прямое взаимодействие с системой и наблюдение за признаками или ссылками на внешние источники данных.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Withhold public documentation that identifies RAG data sources, indexes, and retrieval architecture.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0002 Разведка</span><p>Исследователи Zenity установили, что Microsoft Copilot for M365 индексирует все письма, полученные во входящий ящик, даже если получатель их не открывает.</p></a>
</div>
