---
atlas_id: AML.T0071
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут добавлять поддельные записи в базу данных RAG (генерации, дополненной извлечением) жертвы. Для этого в источник данных, который поступает в базу RAG, помещается содержимое, оформленное так, чтобы...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: False RAG Entry Injection
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0007
title: Внедрение ложной записи RAG
url: /techniques/AML.T0071/
---

Злоумышленники могут добавлять поддельные записи в базу данных RAG (генерации, дополненной извлечением) жертвы. Для этого в источник данных, который поступает в базу RAG, помещается содержимое, оформленное так, чтобы большая языковая модель (LLM) RAG-системы восприняла его как документ. Когда такая RAG-запись извлекается, LLM обманывают так, что часть полученного содержимого принимается за поддельный результат RAG.

Размещение поддельного RAG-документа внутри обычной RAG-записи позволяет обойти инструменты мониторинга данных. Кроме того, такой документ нельзя удалить напрямую.

Злоумышленник может использовать выявленные системные ключевые слова, чтобы понять, как заставить конкретную LLM воспринимать содержимое как RAG-запись. Он может манипулировать метаданными внедренной записи, включая название документа, автора и дату создания.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0007/"><span class="relation-id">AML.TA0007</span><strong>Уклонение от защиты</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0007 Уклонение от защиты</span><p>Когда пользователь ищет банковские реквизиты и извлекается отравленная RAG-запись, маркер `Actual Snippet:` заставляет LLM воспринимать извлеченный текст как фрагмент реального документа.</p></a>
</div>
